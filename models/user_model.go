package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"password" bson:"password" validate:"required,min=6"`
	Email         *string            `json:"email" bson:"email" validate:"email,required"`
	Token         *string            `json:"token" bson:"token"`
	Refresh_token *string            `json:"refresh_token" bson:"refresh_token"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
	Type          string             `json:"type" bson:"type"`
}
