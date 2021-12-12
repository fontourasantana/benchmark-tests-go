package app

import (
	"os"
	"fmt"
	"time"
	"context"
	"net/http"
	"ameicosmeticos/infrastructure/persistence"
	cc "ameicosmeticos/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

type App struct {
	shutdownTimeout time.Duration
}

func Create() *App {
	return &App{
		shutdownTimeout: 5 * time.Second, // Default shutdown timeout
		// Adicionar logger
	}
}

func TimeOutMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 1 * time.Second)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				// write response and abort the request
				// c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.AbortWithStatus(http.StatusGatewayTimeout)
				println("aborted")
			}

			//cancel to clear resources after finished
			cancel()
		}()
		
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	  }
}

func LoadSessions() gin.HandlerFunc {
	store, _ := redis.NewStoreWithPool(persistence.GetCacheConnectionPool(), []byte("amei-secret"))
	return sessions.Sessions("session-id", store)
}

func (app *App) Run() {
	fmt.Println("> [ app ] Start running ...")
	fmt.Println("> [ app ] Creating database connection config")
	dbConnectionConfig := persistence.DBConnectionConfig{
		Host:				os.Getenv("DB_HOST"),
		User:				os.Getenv("DB_USERNAME"),
		Password:			os.Getenv("DB_PASSWORD"),
		Database:			os.Getenv("DB_DATABASE"),
		Port:				os.Getenv("DB_PORT"),
		ConnMaxLifetime:	5 * time.Minute,
		MaxIdleConns:		50,
		MaxOpenConns: 		50,
	}

	fmt.Println("> [ app ] Creating cache connection config")
	cacheConnectionConfig := persistence.CacheConnectionConfig{
		Host:				os.Getenv("REDIS_HOST"),
		Port:				os.Getenv("REDIS_PORT"),
		ConnMaxLifetime:	5 * time.Minute,
		MaxIdleConns:		50,
		MaxOpenConns: 		50,
	}

	persistenceHandler := persistence.NewPersistenceHandler(dbConnectionConfig, cacheConnectionConfig)
	repositoryContainer := NewRepositoryContainer(persistenceHandler)
	serviceContainer := NewServiceContainer(repositoryContainer)
	controllers := NewControllers(serviceContainer)
	engine := LoadRouters(controllers).InitRouter()
	engine.Use(LoadSessions())
	engine.Use(TimeOutMiddleware())

	engine.GET("/incr", func(c *gin.Context) {		
		time.Sleep(time.Second * 2)
		session := sessions.Default(c)
		var count int
		v := session.Get("count")

		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		session.Set("count", count)
		session.Save()
		cc.Response(c, gin.H{"count": count}, 200)// c.JSON(200, gin.H{"count": count})
	})

	fmt.Println("> [ app ] Started successfully")
	// fmt.Println(engine.Routes())	

	engine.Run(":8080")

	// stop := func() {
	// 	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	// 	defer cancel()

	// 	app.logger.Info(ctxWithTimeout, "shutting down...\n")

	// 	var errorOccurred bool
	// 	errCh := make(chan error, len(app.adapters))

	// 	for _, adapter := range app.adapters {
	// 		go func(adapter Adapter) {
	// 			errCh <- adapter.Stop(ctxWithTimeout)
	// 		}(adapter)
	// 	}

	// 	for i := 0; i < len(app.adapters); i++ {
	// 		if err := <-errCh; err != nil {
	// 			errorOccurred = true
	// 			app.logger.Critical(ctxWithTimeout, "shutdown error: %v\n", err)
	// 		}
	// 	}

	// 	app.logger.Info(ctxWithTimeout, "gracefully stopped\n")

	// 	if errorOccurred {
	// 		os.Exit(1)
	// 	}
	// }

	// for _, adapter := range app.adapters {
	// 	go func(adapter Adapter) {
	// 		app.logger.Critical(ctx, "adapter start error: %v\n", adapter.Start(ctx))
	// 	}(adapter)
	// }

	// shutdown.GracefulStop(stop)
}