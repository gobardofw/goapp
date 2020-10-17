package bootstrap

import "github.com/gobardofw/cache"

func init() {
	// {{if eq .cache "redis"}}
	conf := app.Config()
	if c := cache.NewRedisCache(
		"// {{.name}}",
		conf.String("redis.host", "127.0.0.1"),
		conf.Int("redis.maxIdle", 50),
		conf.Int("redis.maxActive", 10000),
		conf.UInt8("redis.cache_db", 1),
	); c != nil {
		_container.Register("--APP-CACHE", c)
	} else {
		panic("failed to build cache driver")
	}
	// {{else}}
	if c := cache.NewFileCache("// {{.name}}", "./storage/cache"); c != nil {
		_container.Register("--APP-CACHE", c)
	} else {
		panic("failed to build cache driver")
	}
	// {{end}}
}
