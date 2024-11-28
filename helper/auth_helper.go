package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.coom/Uttkarsh-raj/RBAC/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Email      string
	First_Name string
	Last_Name  string
	Type       string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(userPass, hashPass string) (bool, string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(userPass))
	if err != nil {
		return false, "Incorrect email or Password.", err
	}
	return true, "", nil
}

func GenerateTokens(email, fname, lname, tpe string) (string, string, error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: fname,
		Last_Name:  lname,
		Type:       tpe,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(48)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic("Error signing token: ", err)
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic("Error signing refresh token: ", err)
		return "", "", err
	}
	return token, refreshToken, nil
}

func UpdateTokens(signedToken, signedRefreshToken, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"email": email}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateObj},
	}, &opt)
	if err != nil {
		log.Panic(err)
		return
	}
}

func VerifyToken(tokenString string, allowedRoles []string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}

	claim, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	for _, val := range allowedRoles {
		if claim.Type == val {
			return token, nil
		}
	}
	return nil, fmt.Errorf("permission denied !! You don't have permissions to access/modify this data")
}
