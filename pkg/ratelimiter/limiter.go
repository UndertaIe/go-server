package ratelimiter

import (
	"context"
	"time"
)

// 实现不同的限流策略，如路由, IP, 资源
type Limiter interface {
	// key生成策略
	Key(c context.Context) string
	// 获取桶
	GetBucket(key string) (Bucket, bool)
	// 添加桶
	AddBuckets(opt ...BucketOption) Limiter
}

// 单一实体限流模型
type Bucket interface {
	// 可用token数量
	Available() int
	// 桶最大容量
	Capacity() int
	// 每秒生成数量
	Rate() float64
	// 获取n个token
	Take(n int) bool
	// 等待获取n个token，超时失败
	Wait(n int, d time.Duration) bool

	Idle() bool

	Reset(opt BucketOption) Bucket
}

// 配置信息:动态修改
type BucketOption struct {
	// 映射bucket
	Key string
	// 时间间隔
	fillInterval time.Duration
	// 容量
	capacity int
	// 增量
	quantum int
	// 空闲间隔
	idleInterval time.Duration
}

var defaultBucketOption = BucketOption{
	fillInterval: time.Minute,
	capacity:     1000,
	quantum:      100,
	idleInterval: time.Minute * 10,
}
