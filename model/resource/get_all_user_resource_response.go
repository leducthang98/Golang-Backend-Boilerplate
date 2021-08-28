package resource

import "boilerplate/model/entity"

type GetAllUserResourceResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func NewGetAllUserResourceResponse(entity *entity.User) *GetAllUserResourceResponse {
	return &GetAllUserResourceResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}
}
