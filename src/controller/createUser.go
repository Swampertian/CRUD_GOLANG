package controller

import (
	"net/http"

	"github.com/Swampertian/CRUD_GOLANG/src/configuration/logger"
	"github.com/Swampertian/CRUD_GOLANG/src/configuration/validation"
	"github.com/Swampertian/CRUD_GOLANG/src/controller/model/request"
	"github.com/Swampertian/CRUD_GOLANG/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error binding JSON: ", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserErro(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    "123456789",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User criado com sucesso",
		zap.String("journey", "createUser"),
	)
	c.JSON(http.StatusCreated, userResponse)
}
