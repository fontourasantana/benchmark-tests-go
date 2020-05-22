package app

import (
	"sync"
	"ameicosmeticos/app/controllers"
)

type IControllers interface {
	InjectCanceledUsersController() *controllers.CanceledUsersController
}

type Controllers struct {
	serviceContainer	IServiceContainer
}

var (
	c		*Controllers
	controllerOnce	sync.Once
)

func NewControllers(serviceContainer IServiceContainer) IControllers {
	if c == nil {
		controllerOnce.Do(func() {
			println("> [ controllers ] Creating ...")
			c = &Controllers{serviceContainer}
			println("> [ controllers ] Created")
		})
	}

	return c
}

func (this *Controllers) InjectCanceledUsersController() *controllers.CanceledUsersController {
	canceledUsersService := this.serviceContainer.GetCanceledUsersService()
	return &controllers.CanceledUsersController{canceledUsersService}
}
