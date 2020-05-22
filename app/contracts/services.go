package contracts

import "ameicosmeticos/app/models"

type ICanceledUsersService interface {
	GetUserById(uint64) (models.CanceledUsers, error)
}