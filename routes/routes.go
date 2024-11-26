package routes

import (
	"github.com/gin-gonic/gin"
	"github.coom/Uttkarsh-raj/RBAC/controller"
	"github.coom/Uttkarsh-raj/RBAC/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(middleware.CORSMiddleware())
	server.POST("/signup", controller.SignUpUser())
	server.POST("/login", controller.LogInUser())
}

func RegisterGenericRoutes(server *gin.Engine) {
	allowed := []string{"User", "Moderator", "Admin"}
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.CheckAuthAndPermissions(allowed))
	server.GET("/users", controller.GetAllUsers())
	server.GET("/users/:id", controller.GetUser())
}

func RegisterModeratorRoutes(server *gin.Engine) {
	allowed := []string{"Admin", "Moderator"}
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.CheckAuthAndPermissions(allowed))
	server.PATCH("/update/:id", controller.UpdateUser())
}

func RegisterAdminRoutes(server *gin.Engine) {
	allowed := []string{"Admin"}
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.CheckAuthAndPermissions(allowed))
	server.DELETE("/delete/:id", controller.DeleteUser())
}
