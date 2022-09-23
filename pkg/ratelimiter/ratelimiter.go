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
	if cfg.PoolEnabled() {
		cl.pool = &sync.Pool{
			New: func() any {
				return newBucketer(cfg.BucketOption)
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
	BucketOption
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
func (lo *LimiterOption) DefaultBucketOption() BucketOption {
	return lo.BucketOption
}
func (lo *LimiterOption) FillInterval() time.Duration {
	return lo.fillInterval
}
func (lo *LimiterOption) Capacity() int {
	return lo.capacity
}
func (lo *LimiterOption) Quantum() int {
	return lo.quantum
}
func (lo *LimiterOption) IdleInterval() time.Duration {
	return lo.idleInterval
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
	if lo.fillInterval == 0 {
		lo.fillInterval = defaultBucketOption.fillInterval
	}
	if lo.capacity == 0 {
		lo.capacity = defaultBucketOption.capacity
	}
	if lo.quantum == 0 {
		lo.quantum = defaultBucketOption.quantum
	}
	if lo.idleInterval == 0 {
		lo.idleInterval = defaultBucketOption.idleInterval
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
	lo.BucketOption = bo
	return lo
}

func WithBucketOption(bo BucketOption) Option {
	return bo
}

func (cl *rateLimiter) Key(c context.Context) string {
	return cl.cfg.keyFunc(c)
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
	opt = cl.ensureBucketOption(opt)
	if cl.cfg.PoolEnabled() {
		cl.addBucketWithPool(opt)
		return
	}
	cl.buckets[opt.Key] = newBucketer(opt)
}

func (cl *rateLimiter) ensureBucketOption(opt BucketOption) BucketOption {
	if opt.fillInterval == 0 {
		opt.fillInterval = cl.cfg.FillInterval()
	}
	if opt.capacity == 0 {
		opt.capacity = cl.cfg.Capacity()
	}
	if opt.quantum == 0 {
		opt.quantum = cl.cfg.Quantum()
	}
	if opt.idleInterval == 0 {
		opt.idleInterval = cl.cfg.IdleInterval()
	}
	return opt
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

func NewBucketOption() BucketOption {
	return BucketOption{}
}

func (bo BucketOption) WithBucketFillInterval(fillInterval time.Duration) BucketOption {
	bo.fillInterval = fillInterval
	return bo
}

func (bo BucketOption) WithBucketcCapacity(capacity int) BucketOption {
	bo.capacity = capacity
	return bo
}

func (bo BucketOption) WithBucketQuantum(quantum int) BucketOption {
	bo.quantum = quantum
	return bo
}

func (bo BucketOption) WithBucketIdleInterval(idleInterval time.Duration) BucketOption {
	bo.idleInterval = idleInterval
	return bo
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
