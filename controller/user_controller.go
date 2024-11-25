package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.coom/Uttkarsh-raj/RBAC/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*4)
		defer cancel()
		cursor, err := userCollection.Find(c, bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		defer cursor.Close(c)

		var users []models.UserResponse
		for cursor.Next(c) {
			var user models.User
			if err := cursor.Decode(&user); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			filteredUser := models.UserResponse{
				ID:        user.ID,
				Email:     *user.Email,
				FirstName: *user.First_name,
				LastName:  *user.Last_name,
			}
			users = append(users, filteredUser)
		}

		if err := cursor.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"users": users}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		id := ctx.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID format"})
			return
		}

		var user models.UserResponse
		cur := userCollection.FindOne(c, bson.M{"_id": objectID})
		err = cur.Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"user": user}})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*8)
		defer cancel()
		id := ctx.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID format"})
			return
		}

		var updates map[string]interface{}
		if err = json.NewDecoder(ctx.Request.Body).Decode(&updates); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid request body"})
			return
		}

		var existingUser models.UserResponse
		if err = userCollection.FindOne(c, bson.M{"_id": objectID}).Decode(&existingUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "User not found"})
			return
		}
		delete(updates, "_id")

		if email, ok := updates["email"].(string); ok && email != "" {
			count, err := userCollection.CountDocuments(c, bson.M{"email": email})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}
			if count > 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Email already exists"})
				return
			}
		}

		update := bson.M{"$set": updates}

		_, err = userCollection.UpdateByID(c, objectID, update)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "User updated successfully"})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		id := ctx.Param("id")
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid ID format"})
			return
		}
		var user models.UserResponse
		cur := userCollection.FindOneAndDelete(c, bson.M{"_id": objectID})
		err = cur.Decode(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"user": user}})
	}
}
