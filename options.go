package scheduler

import "time"

//在Spring中这时需要设置concurrent的值为false, 禁止并发执行。<property name="concurrent" value="true" />
const (
	defaultStart      = "00:00:00"
	defaultInterval   = time.Second
	defaultConcurrent = true
)

type Config struct {
	start      string        //开始时间
	interval   time.Duration //间隔时间
	concurrent bool          //禁止并发
}

type Option func(*Config)

func Start(start string) Option {
	return func(c *Config) {
		c.start = start
	}
}

func Interval(interval time.Duration) Option {
	return func(c *Config) {
		c.interval = interval
	}
}

func Concurrent(concurrent bool) Option {
	return func(c *Config) {
		c.concurrent = concurrent
	}
}
