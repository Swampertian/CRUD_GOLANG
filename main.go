package main

import (
	"context"
	"github.com/Swampertian/CRUD_GOLANG/src/configuration/logger"
	"github.com/Swampertian/CRUD_GOLANG/src/configuration/mongodb"
	"github.com/Swampertian/CRUD_GOLANG/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, UserController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
