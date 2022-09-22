package ratelimiter

//TODO:  限流策略链满足所有Limiter则通过限流策略

type Policy int

const (
	AllPass Policy = iota
	AnyPass
	HalfPass
)

type LimiterChain []Limiter
