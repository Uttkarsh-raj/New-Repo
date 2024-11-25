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
	routes.RegisterRoutes(server)
	log.Fatal(server.Run(":3000"))
}
