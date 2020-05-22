package repository

import "ameicosmeticos/domains/administrative/entity"

type CanceledUsersRepository interface {
	GetAll() ([]entity.CanceledUsers, error)
	GetUser(uint64) (*entity.CanceledUsers, error)
}