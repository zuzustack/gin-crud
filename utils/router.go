package utils

import (
	"zuzustack/learn/api/controllers"

	"github.com/gin-gonic/gin"
)


func SetRouter(r *gin.Engine){
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}