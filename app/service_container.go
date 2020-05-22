package app

import (
	"sync"
	"ameicosmeticos/app/contracts"
	"ameicosmeticos/app/services"
)

type IServiceContainer interface {
	GetCanceledUsersService() contracts.ICanceledUsersService
}

type ServiceContainer struct {
	repositoryContainer	IRepositoryContainer
}

var (
	serviceContainer	*ServiceContainer
	servContainerOnce	sync.Once
)

func NewServiceContainer(repositoryContainer IRepositoryContainer) IServiceContainer {
	if serviceContainer == nil {
		servContainerOnce.Do(func() {
			println("> [ service container ] Creating ...")
			serviceContainer = &ServiceContainer{repositoryContainer}
			println("> [ service container ] Created")
		})
	}

	return serviceContainer
}

func (this *ServiceContainer) GetCanceledUsersService() contracts.ICanceledUsersService {
	canceledUsersRepository := this.repositoryContainer.GetCanceledUsersRepository()	
	return services.NewCanceledUsersService(canceledUsersRepository)
}