package persistence

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"log"
)

type CacheHandler struct {
	Conn *redis.Pool
}

var (
	cache		*CacheHandler
	cacheOnce	sync.Once
)

const (
	redisExpire = 60
)

func NewCacheHandler(config CacheConnectionConfig) *CacheHandler {
	if cache == nil {
		cacheOnce.Do(func() {
			println("> [ infrastructure ] Creating cache handler ...")
			cache = &CacheHandler{
				&redis.Pool{
					MaxIdle: config.MaxIdleConns,
					MaxActive: config.MaxOpenConns,
					Wait: true,
					IdleTimeout: config.ConnMaxLifetime,
					Dial: func() (redis.Conn, error) {
						return redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
					},
				},
			}

			conn := cache.Conn.Get()
			defer conn.Close()

			_, err := conn.Do("PING")
			if err != nil {
				log.SetFlags(0)
				log.Fatal("> [ infrastructure ] Não foi possivel estabelecer conexão com o CACHE")
			}

			println("> [ infrastructure ] Cache handler created")
		})
	}

	return cache
}

func (this *CacheHandler) Set(key string, value []byte) error {
	conn := this.Conn.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", key, redisExpire, value)
	return err
}

func (this *CacheHandler) Get(key string) ([]byte, error) {
	conn := this.Conn.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func (this *CacheHandler) Delete(key string) error {
	conn := this.Conn.Get()
	defer conn.Close()
	_, err := redis.Bytes(conn.Do("DEL", key))
	return err
}
