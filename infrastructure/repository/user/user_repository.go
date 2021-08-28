package user

import (
	"boilerplate/infrastructure/db/rdb"
	"boilerplate/model/entity"
)

type UserRepository struct {
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetAll() (entities []*entity.User, err error) {
	err = rdb.GetInstance().Raw("Select * from `user`").Scan(&entities).Error
	return
}
