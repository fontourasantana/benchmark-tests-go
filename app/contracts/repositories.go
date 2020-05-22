package contracts

import "ameicosmeticos/app/models"

type ICanceledUsersRepository interface {
	GetUser(uint64) (models.CanceledUsers, error)
}