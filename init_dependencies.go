package main

import (
	"github.com/Swampertian/CRUD_GOLANG/src/controller"
	"github.com/Swampertian/CRUD_GOLANG/src/model/repository"
	"github.com/Swampertian/CRUD_GOLANG/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
