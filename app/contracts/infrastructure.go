package contracts

import "github.com/gomodule/redigo/redis"

type IPersistenceHandler interface {
	GetDBHandler() IDbHandler
	GetCacheHandler() ICacheHandler
	GetCacheConnectionPool() *redis.Pool
}

type IDbHandler interface {
	Execute(statement string)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error	
}

type ICacheHandler interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}