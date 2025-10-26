package converter

import (
	"github.com/Swampertian/CRUD_GOLANG/src/model"
	"github.com/Swampertian/CRUD_GOLANG/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
