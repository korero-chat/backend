package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Username     string             `json:"username"`
	PasswordHash string             `json:"passwordhash"`
	Email        string             `json:"email"`
}
