package models

type User struct {
	UserID   string `bson:"_id, omitempty", json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
