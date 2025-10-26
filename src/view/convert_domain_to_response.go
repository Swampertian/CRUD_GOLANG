package view

import (
	"github.com/Swampertian/CRUD_GOLANG/src/controller/model/response"
	"github.com/Swampertian/CRUD_GOLANG/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
