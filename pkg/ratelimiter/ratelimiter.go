package ratelimiter

import (
	"context"
	"sync"
	"time"

	"github.com/juju/ratelimit"
)

var _ Limiter = (*rateLimiter)(nil)

type sig struct{}

type rateLimiter struct {
	buckets     map[string]Bucket
	pool        *sync.Pool
	keyFunc     KeyFunc
	stopRefresh chan sig
	mux         sync.RWMutex
	cfg         LimiterOption
}

func NewRateLimiter(options ...Option) *rateLimiter {
	cfg := LimiterOption{}
	for _, opt := range options {
		cfg = opt.apply(cfg)
	}
	cfg.ensureInit()
	cl := &rateLimiter{
		buckets:     make(map[string]Bucket),
		stopRefresh: make(chan sig, 1),
		cfg:         cfg,
	}
	if cfg.poolEnabled {
		cl.pool = &sync.Pool{
			New: func() any {
				return newBucketer(*cfg.defaultBucketOption)
			},
		}
	}
	cl.innerStart()
	return cl
}

type LimiterOption struct {
	// key生成策略
	keyFunc KeyFunc
	// 开启回收Bucket
	poolEnabled bool
	// 开启定时刷新 删除回收不用的bucket
	refreshEnabled bool
	// 定时刷新间隔
	refreshInterval time.Duration
	// 默认bucket配置
	defaultBucketOption *BucketOption
}

func (lo *LimiterOption) PoolEnabled() bool {
	return lo.poolEnabled
}
func (lo *LimiterOption) RefreshEnabled() bool {
	return lo.refreshEnabled
}
func (lo *LimiterOption) RefreshInterval() time.Duration {
	return lo.refreshInterval
}
func (lo *LimiterOption) DefaultBucketOption() *BucketOption {
	return lo.defaultBucketOption
}

// 默认刷新时间间隔
const defaultRefreshInterval = time.Minute * 10

func (lo *LimiterOption) ensureInit() {
	if lo.keyFunc == nil {
		panic("need call ratelimiter.WithKeyFunc()")
	}
	if lo.refreshInterval == 0 {
		lo.refreshInterval = defaultRefreshInterval
	}
	bopt := lo.defaultBucketOption
	if bopt.fillInterval == 0 {
		bopt.fillInterval = defaultBucketOption.fillInterval
	}
	if bopt.capacity == 0 {
		bopt.capacity = defaultBucketOption.capacity
	}
	if bopt.quantum == 0 {
		bopt.quantum = defaultBucketOption.quantum
	}
	if bopt.idleInterval == 0 {
		bopt.idleInterval = defaultBucketOption.idleInterval
	}
}

type Option interface {
	apply(LimiterOption) LimiterOption
}

type refreshInterval time.Duration

func (ri refreshInterval) apply(lo LimiterOption) LimiterOption {
	lo.refreshEnabled = true
	lo.refreshInterval = time.Duration(ri)
	return lo
}

func WithRefreshInterval(d time.Duration) Option {
	return refreshInterval(d)
}

type poolEnabled bool

func (pe poolEnabled) apply(lo LimiterOption) LimiterOption {
	lo.poolEnabled = bool(pe)
	return lo
}

func WithPool() Option {
	return poolEnabled(true)
}

func (bo BucketOption) apply(lo LimiterOption) LimiterOption {
	lo.defaultBucketOption = &bo
	return lo
}

type KeyFunc func(c context.Context) string

func (f KeyFunc) apply(lo LimiterOption) LimiterOption {
	lo.keyFunc = f
	return lo
}

func WithKeyFunc(f func(c context.Context) string) Option {
	return KeyFunc(f)
}

func WithBucketOption(bo BucketOption) Option {
	return bo
}

func (cl *rateLimiter) Key(c context.Context) string {
	return cl.keyFunc(c)
}

func (cl *rateLimiter) GetBucket(key string) (Bucket, bool) {
	cl.mux.RLock()
	defer cl.mux.RUnlock()
	bucket, ok := cl.buckets[key]
	return bucket, ok
}

func (cl *rateLimiter) AddBuckets(opts ...BucketOption) Limiter {
	cl.mux.Lock()
	defer cl.mux.Unlock()
	for _, opt := range opts {
		if _, ok := cl.buckets[opt.Key]; !ok {
			cl.addBucket(opt)
		}
	}
	return cl
}

func (cl *rateLimiter) addBucketWithPool(opt BucketOption) {
	rb, _ := cl.pool.Get().(Bucket)
	cl.buckets[opt.Key] = rb.Reset(opt)
}

func (cl *rateLimiter) addBucket(opt BucketOption) {
	if cl.cfg.PoolEnabled() {
		cl.addBucketWithPool(opt)
		return
	}
	cl.buckets[opt.Key] = newBucketer(opt)
}

func (cl *rateLimiter) ShutDown() bool {
	cl.stopRefresh <- sig{}
	return true
}

func (cl *rateLimiter) innerStart() {
	if cl.cfg.RefreshEnabled() {
		go cl.refresh()
	}
}

func (cl *rateLimiter) refresh() {
	tk := time.NewTicker(cl.cfg.RefreshInterval())
	for range tk.C {
		select {
		case <-cl.stopRefresh: // 接收到shutdown信号退出
			close(cl.stopRefresh)
			return
		default:
		}
		cl.mux.Lock()
		var idlers []string
		for k, v := range cl.buckets {
			if v.Idle() {
				idlers = append(idlers, k)
				if cl.cfg.PoolEnabled() {
					cl.pool.Put(v)
				}
			}
		}
		for _, key := range idlers {
			delete(cl.buckets, key)
		}
		cl.mux.Unlock()
	}
}

var _ Bucket = (*Bucketer)(nil)

type Bucketer struct {
	*ratelimit.Bucket
	idleInterval time.Duration
}

func newBucketer(opt BucketOption) *Bucketer {
	bucket := ratelimit.NewBucketWithQuantum(
		opt.fillInterval,
		int64(opt.capacity),
		int64(opt.quantum),
	)
	return &Bucketer{
		Bucket:       bucket,
		idleInterval: opt.idleInterval}
}

func (rb *Bucketer) Available() int {
	return int(rb.Bucket.Available())
}

func (rb *Bucketer) Capacity() int {
	return int(rb.Bucket.Capacity())
}

func (rb *Bucketer) Rate() float64 {
	return rb.Bucket.Rate()
}

func (rb *Bucketer) Take(n int) bool {
	nn := int64(n)
	return rb.Bucket.TakeAvailable(nn) == nn
}

const (
	Forever        = time.Duration(0x7fffffffffffffff)
	DefaultTimeout = time.Second * 60
)

func (rb *Bucketer) Wait(n int, timeout time.Duration) bool {
	return rb.Bucket.WaitMaxDuration(int64(n), timeout)
}

func (rb *Bucketer) Idle() bool {
	return rb.Bucket.IdleTime() > rb.idleInterval
}

func (rb *Bucketer) Reset(opt BucketOption) Bucket {
	rb.Bucket.Reset(opt.fillInterval, int64(opt.capacity), int64(opt.quantum))
	rb.idleInterval = opt.idleInterval
	return rb
}
