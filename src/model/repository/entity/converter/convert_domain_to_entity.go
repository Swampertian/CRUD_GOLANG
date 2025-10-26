package converter

import (
	"github.com/Swampertian/CRUD_GOLANG/src/model"
	"github.com/Swampertian/CRUD_GOLANG/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
