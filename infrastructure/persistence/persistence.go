package persistence

import (
	"sync"
	"time"
	"ameicosmeticos/app/contracts"
	"github.com/gomodule/redigo/redis"
)


type DBConnectionConfig struct {
	Host            string
	User            string
	Password        string
	Database        string
	Port            string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

type CacheConnectionConfig struct {
	Host			string
	Port			string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}


type PersistenceHandler struct {
	dbHandler		contracts.IDbHandler
	cacheHandler	contracts.ICacheHandler
}

var (
	persistenceHandler	*PersistenceHandler
	persistenceOnce		sync.Once
)

func NewPersistenceHandler(dbConfig DBConnectionConfig, cacheConfig CacheConnectionConfig) contracts.IPersistenceHandler {
	if persistenceHandler == nil {
		persistenceOnce.Do(func() {
			println("> [ infrastructure ] Creating persistence handler ...")
			persistenceHandler = &PersistenceHandler{
				NewDBHandler(dbConfig),
				NewCacheHandler(cacheConfig),
			}
			println("> [ infrastructure ] Persistence handler created")
		})
	}

	return persistenceHandler
}

func GetCacheConnectionPool() *redis.Pool {
	return persistenceHandler.GetCacheConnectionPool()
}

func (this *PersistenceHandler) GetDBHandler() contracts.IDbHandler {
	return this.dbHandler
}

func (this *PersistenceHandler) GetCacheHandler() contracts.ICacheHandler {
	return this.cacheHandler
}

func (this *PersistenceHandler) GetCacheConnectionPool() *redis.Pool {
	var iCacheHandler interface{} = this.cacheHandler
	cacheHandler, _ := iCacheHandler.(*CacheHandler)

	return cacheHandler.Conn
}