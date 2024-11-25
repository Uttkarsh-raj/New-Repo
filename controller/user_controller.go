package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// for admins and moderators
		ctx.JSON(http.StatusOK, gin.H{"message": "Needs Implementation"})
	}
}
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// for admin and moderators
		ctx.JSON(http.StatusOK, gin.H{"message": "Needs Implementation"})
	}
}
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// for admin and moderators
		ctx.JSON(http.StatusOK, gin.H{"message": "Needs Implementation"})
	}
}
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// for admin
		ctx.JSON(http.StatusOK, gin.H{"message": "Needs Implementation"})
	}
}
