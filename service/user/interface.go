package user

import "boilerplate/model/resource"

type IUserService interface {
	GetAll() ([]*resource.GetAllUserResourceResponse, error)
}
