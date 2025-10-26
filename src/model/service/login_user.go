package service

import (
	"github.com/Swampertian/CRUD_GOLANG/src/configuration/logger"
	"github.com/Swampertian/CRUD_GOLANG/src/configuration/rest_err"
	"github.com/Swampertian/CRUD_GOLANG/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {

	logger.Info("Init loginUser model.",
		zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))
	return user, token, nil
}
