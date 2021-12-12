package app

import (
	"sync"
	"ameicosmeticos/app/contracts"
	"ameicosmeticos/app/repositories"
)

type IRepositoryContainer interface {
	GetCanceledUsersRepository() contracts.ICanceledUsersRepository
}

type RepositoryContainer struct {
	dbHandler		contracts.IDbHandler
	cacheHandler	contracts.ICacheHandler
}

var (
	repositoryContainer		*RepositoryContainer
	repoContainerOnce		sync.Once
)

func NewRepositoryContainer(persistence contracts.IPersistenceHandler) IRepositoryContainer {
	if repositoryContainer == nil {
		repoContainerOnce.Do(func() {
			println("> [ repository container ] Creating ...")
			repositoryContainer = &RepositoryContainer{
				persistence.GetDBHandler(),
				persistence.GetCacheHandler(),
			}
			println("> [ repository container ] Created")
		})
	}

	return repositoryContainer
}

func (this *RepositoryContainer) GetCanceledUsersRepository() contracts.ICanceledUsersRepository {
	return repositories.NewCanceledUsersRepository(this.dbHandler, this.cacheHandler)
}