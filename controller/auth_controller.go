package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/korero-chat/backend/database"
	"github.com/korero-chat/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var response models.ResponseModel

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//validate username and email
	if user.Username == "" {
		response.Error = "Validation error, blank username"
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Email == "" {
		response.Error = "Validation error, blank email"
		json.NewEncoder(w).Encode(response)
	}

	_, err = database.FindUserByUsername(user.Username)
	if err != nil {
		// If username is not found, hash password
		if err.Error() == "mongo: no documents in result" {
			passwordhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				response.Error = "Error while hashing the password"
				json.NewEncoder(w).Encode(response)
			}

			//switch clean password with hashed
			user.Password = string(passwordhash)

			//Insert new user into DB
			err = database.InsertUser(user)
			if err != nil {
				response.Error = err.Error()
				json.NewEncoder(w).Encode(response)
				return
			}

			//Registration successfull
			response.Result = "Registration Successfull"
			json.NewEncoder(w).Encode(response)
			return
		}
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)

	}

	response.Result = "Username already Exists!"
	json.NewEncoder(w).Encode(response)
	return
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {

}

func LogoutEndpoint(w http.ResponseWriter, r *http.Request) {

}
