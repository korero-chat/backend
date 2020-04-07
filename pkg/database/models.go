package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id, omitempty", json:"user_id"`
	Username       string             `json:"username"`
	PasswordHash   string             `json:"passwordhash"`
	Email          string             `json:"email"`
	EmailConfirmed bool               `json:"email_confirmed"`
	Chats          []Chat             `json:"chats"`
}

type Chat struct {
	ID              primitive.ObjectID `bson:"_id, omitempty", json:"id"`
	ChatName        string             `json:"chat_name"`
	ChatCreatorID   string             `json:"chat_creator_id"`
	Members         []User             `json:"chat_members"`
	RoomMembersSize int                `json:"room_members_size"`
}
