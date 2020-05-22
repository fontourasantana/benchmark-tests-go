package services

import (
	"ameicosmeticos/app/contracts"
	"ameicosmeticos/app/models"
)

type CanceledUsersService struct {
	contracts.ICanceledUsersRepository
}

func NewCanceledUsersService(repository contracts.ICanceledUsersRepository) *CanceledUsersService {
	return &CanceledUsersService{repository}
}

func (this *CanceledUsersService) GetUserById(id uint64) (models.CanceledUsers, error) {
	return this.GetUser(id)
}