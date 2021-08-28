package user

import "boilerplate/model/entity"

type IUserRepository interface {
	GetAll() ([]*entity.User, error)
}
