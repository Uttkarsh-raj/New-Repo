package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.coom/Uttkarsh-raj/RBAC/helper"
)

func CheckAuthAndPermissions(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Token missing in the header."})
			c.AbortWithStatus(400)
			return
		}
		// fmt.Println(tokenString[7:])
		_, err := helper.VerifyToken(tokenString[7:], allowedRoles)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error verifying the token: " + err.Error()})
			c.AbortWithStatus(400)
			return
		}
		c.Next()
	}
}
