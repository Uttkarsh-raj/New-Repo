package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.coom/Uttkarsh-raj/RBAC/database"
	"github.coom/Uttkarsh-raj/RBAC/helper"
	"github.coom/Uttkarsh-raj/RBAC/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
		defer cancel()

		var user models.User
		err := json.NewDecoder(c.Request.Body).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": validationErr.Error()})
			return
		}

		userCount, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured while checking for the email", "success": false})
			return
		}
		if userCount > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "this email or phone number already exists", "success": false})
			return
		}

		password, err := helper.HashPassword(*user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "success": false})
			return
		}
		user.Password = &password

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		token, refreshToken, _ := helper.GenerateTokens(*user.Email, *user.First_name, *user.Last_name, user.Type)

		user.Token = &token
		user.Refresh_token = &refreshToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := "User item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "success": false})
			return
		}
		defer cancel()
		*user.Password = ""

		response := gin.H{
			"InsertedID": resultInsertionNumber.InsertedID,
			"user":       user,
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "data": response})
	}
}

func LogInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()

		var user models.User
		var foundUser models.User
		err := json.NewDecoder(c.Request.Body).Decode(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
			return
		}

		err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "login or password is incorrect", "success": false})
			return
		}

		passwordIsValid, msg, err := helper.VerifyPassword(*user.Password, *foundUser.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "success": false})
			return
		}
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"message": msg, "success": false})
			return
		}

		token, refreshToken, err := helper.GenerateTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error generating tokens", "success": false})
			return
		}

		helper.UpdateTokens(token, refreshToken, *foundUser.Email)
		*foundUser.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": gin.H{
				"user": foundUser,
			},
		})

	}
}
