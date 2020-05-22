package service

import (
	"ameicosmeticos/domains/administrative/entity"
)

type CanceledUsersService interface {
	GetAll() ([]entity.CanceledUsers, error)
	// SaveUser(*entity.User) (*entity.User, map[string]string)
	// GetUser(uint64) (*entity.User, error)
	// GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
}