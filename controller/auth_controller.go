package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/goware/emailx"
	"github.com/joho/godotenv"
	"github.com/korero-chat/backend/database"
	"github.com/korero-chat/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var newUser models.NewUserRequest
	var response models.ResponseModel

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &newUser)
	if err != nil {
		w.WriteHeader(500)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//validate username and password

	if errs := validator.Validate(newUser); errs != nil {
		w.WriteHeader(422)
		response.Error = errs.Error()
		response.Result = "Validation error"
		json.NewEncoder(w).Encode(response)
		return
	}

	//check if passwords match
	if newUser.Password != newUser.Password2 {
		w.WriteHeader(422)
		response.Result = "Passwords do not match"
		json.NewEncoder(w).Encode(response)
		return
	}

	//validate email
	err = emailx.Validate(newUser.Email)
	if err != nil {
		if err == emailx.ErrInvalidFormat {
			w.WriteHeader(422)
			response.Error = "Invalid email format"
			json.NewEncoder(w).Encode(response)
			return
		}

		if err == emailx.ErrUnresolvableHost {
			w.WriteHeader(422)
			response.Error = "Unresovlable email host"
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(422)
		response.Error = "Email validation error"
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = database.FindUserByUsername(newUser.Username)
	if err != nil {
		// If username is not found, hash password
		if err.Error() == "mongo: no documents in result" {
			user := models.User{Username: newUser.Username, Email: newUser.Email, Password: newUser.Password}

			passwordhash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				w.WriteHeader(500)
				response.Error = "Error while hashing the password"
				json.NewEncoder(w).Encode(response)
			}

			//switch clean password with hashed
			user.Password = string(passwordhash)

			//Insert new user into DB
			err = database.InsertUser(user)
			if err != nil {
				w.WriteHeader(500)
				response.Error = err.Error()
				json.NewEncoder(w).Encode(response)
				return
			}

			//Registration successfull
			w.WriteHeader(201)
			response.Result = "Registration Successfull"
			json.NewEncoder(w).Encode(response)
			return
		}

		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)

	}
	//username already taken
	w.WriteHeader(409)
	response.Error = "Username already taken"
	json.NewEncoder(w).Encode(response)
	return
}

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user models.User
	var response models.ResponseModel

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(500)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//Check if user with given username exists
	result, err := database.FindUserByUsername(user.Username)
	if err != nil {
		w.WriteHeader(422)
		response.Error = "Invalid username"
		json.NewEncoder(w).Encode(response)
		return
	}

	// Compare given password and hashed
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		w.WriteHeader(401)
		response.Error = "Invalid Password"
		json.NewEncoder(w).Encode(response)
		return
	}

	tk := models.Token{
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte("SECRET_JWT_KEY"))
	if err != nil {
		w.WriteHeader(500)
		response.Error = "Could not create token"
		json.NewEncoder(w).Encode(response)
		return
	}

	var resp = map[string]interface{}{"status": false, "message": "logged_in"}
	resp["user"] = user.Username
	resp["token"] = tokenString

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		response := models.ResponseModel{}
		token := r.Header.Get("x-access-token")
		token = strings.TrimSpace(token)

		if token == "" {
			w.WriteHeader(403)
			response.Error = "Missing auth token"
			json.NewEncoder(w).Encode(response)
		}

		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET_JWT_KEY"), nil
		})
		if err != nil {
			w.WriteHeader(403)
			response.Error = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := context.WithValue(r.Context(), "username", tk)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
