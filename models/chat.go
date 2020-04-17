package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ChatID  primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name    string             `json:"name"`
	Founder User               `json:"founder"`
	Members []User             `json:"members"`
}

type NewChatRequest struct {
	Name    string `validate:"min=3,max=30,regexp=^[a-zA-Z]*$"`
	Founder string `json:"founder"`
}
