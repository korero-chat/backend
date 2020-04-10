package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/korero-chat/backend/database"
	"github.com/korero-chat/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
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

	//validate data
	nur := models.NewUserRequest{Username: user.Username, Password: user.Password, Email: user.Email}
	if errs := validator.Validate(nur); errs != nil {
		response.Error = errs.Error()
		response.Result = "Validation error"
		json.NewEncoder(w).Encode(response)
		return
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

	json.NewEncoder(w).Encode(response)
	return
}

/*func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var response models.ResponseModel

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil{
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//Check if user with given username exists
	result, err = database.FindUserByUsername(user.Username)
	if err != nil{
		response.Error = "Invalid username"
		json.NewEncoder(w).Encode(response)
		return
	}

	// Compare given password and hashed
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		response.Error = "Invalid Password"
		json.NewEncoder(w).Encode(response)
		return
	}

	//Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
		"password": result.Password
	})

	//load secret token key from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT_KEY")))



func LogoutEndpoint(w http.ResponseWriter, r *http.Request) {

}
*/
