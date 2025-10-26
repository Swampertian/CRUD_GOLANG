package routes

import (
	"github.com/Swampertian/CRUD_GOLANG/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userID", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userID", controller.UpdateUser)
	r.DELETE("/deleteUser/:userID", controller.DeleteUser)
}
