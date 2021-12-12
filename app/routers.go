package app

import (
	"sync"
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	InitRouter() *gin.Engine
}

type Router struct {
	IControllers
}

func (this *Router) InitRouter() *gin.Engine {
	println("> [ routers ] Initializing routes ...")
	controller := this.InjectCanceledUsersController()

	routers := gin.Default()
	routers.GET("/", controller.FindById)

	println("> [ routers ] Routes started successfully")
	return routers
}

var (
	router		*Router
	routerOnce	sync.Once
)

func LoadRouters(controllers IControllers) IRouter {
	if router == nil {
		routerOnce.Do(func() {
			println("> [ routers ] Creating router manager ...")
			router = &Router{controllers}
			println("> [ routers ] Router manager created")
		})
	}
	return router
}