package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.coom/Uttkarsh-raj/RBAC/routes"
)

func main() {
	fmt.Println("Starting server...")
	server := gin.New()
	server.Use(gin.Logger())
	routes.RegisterRoutes(server)
	routes.RegisterGenericRoutes(server)
	routes.RegisterModeratorRoutes(server)
	routes.RegisterAdminRoutes(server)
	log.Fatal(server.Run(":3000"))
}
