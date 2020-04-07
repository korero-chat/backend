package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	// Validate username and password / check if blank
	if user.Username == "" || user.PasswordHash == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

func GetUsersChatsEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var username string
	json.NewDecoder(r.Body).Decode(&username)

	result := database.GetChatsByUser(username)
	json.NewEncoder(w).Encode(result)
}

func GetUsersOfChatEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var users []database.User
	params := mux.Vars(r)

	users = database.GetUsersOfChat(params["id"])
	json.NewEncoder(w).Encode(users)
}

func CreateNewChatRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var chat database.Chat
	_ = json.NewDecoder(r.Body).Decode((&chat))
	database.InsertChat(chat)

}
