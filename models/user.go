package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

type NewUserRequest struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Password string `validate:"min=8"`
	Email    string `validate:regexp="^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
}
