package user

import (
	"boilerplate/infrastructure/repository/user"
	"boilerplate/model/resource"
	"boilerplate/pkg/constant"
	"boilerplate/server/container"
)

type UserService struct {
	userRepository *user.UserRepository
}

func NewUserService() IUserService {
	return &UserService{
		userRepository: container.GetInstance().Get(constant.UserIocRepository).(*user.UserRepository),
	}
}

func (u *UserService) GetAll() (listUserResource []*resource.GetAllUserResourceResponse, err error) {
	listUserEntities, err := u.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	for _, userEntity := range listUserEntities {
		listUserResource = append(listUserResource, resource.NewGetAllUserResourceResponse(userEntity))
	}

	return
}
