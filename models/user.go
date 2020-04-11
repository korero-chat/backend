package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewUserRequest struct {
	Username string `validate:"min=3,max=20,regexp=^[a-zA-Z]*$"`
	Password string `validate:"min=8"`
	Email    string `validate:"min=6,regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$"`
}

type Token struct {
	Username string `json:"username"`
	*jwt.StandardClaims
}
