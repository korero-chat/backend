package models

type User struct {
	UserID   string `bson:"_id, omitempty", json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewUserRequest struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Password string `validate:"min=8"`
}
