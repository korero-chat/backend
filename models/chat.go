package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ChatID  primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name    string             `json:"name"`
	Founder User               `json:"founder"`
	Members []User             `json:"members"`
}
