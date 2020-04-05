package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/korero-chat/backend/pkg/crypto"
	database "github.com/korero-chat/backend/pkg/database"
)

func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get data from json request and store them in user variable
	var user database.User
	_ = json.NewDecoder(r.Body).Decode((&user))

	//Check if email already taken
	result := database.CheckIfEmailTaken(user.Email)
	if result == false {
		w.WriteHeader(http.StatusConflict)
		return
	}

	// Hash password
	passwordhash, err := crypto.HashPassword(user.PasswordHash)
	if err != nil {
		log.Fatalf("[-] Error while hashing the password")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Insert user into the database
	database.InsertUser(user.Username, passwordhash, user.Email)

	json.NewEncoder(w).Encode(user)

}

func GetChatsByUserIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func GetChatUsersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var chats []database.Chat

}
