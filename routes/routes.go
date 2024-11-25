package routes

import (
	"github.com/gin-gonic/gin"
	"github.coom/Uttkarsh-raj/RBAC/controller"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", controller.SignUpUser())
	server.POST("/login", controller.LogInUser())
	server.GET("/users", controller.GetAllUsers())
	server.GET("/users/:id", controller.GetUser())
	server.PATCH("/users/:id", controller.UpdateUser())
	server.DELETE("/users/:id", controller.DeleteUser())
}
