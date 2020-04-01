package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/korero-chat/backend/pkg/database"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Struct that will be encoded to a JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	//Get the JSON body and decode into Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If structure of the body is wrong, return HTTP error and retrn
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get expected password from Database
	expectedPassword := database.GetExpectedPassword(creds.Username)

	if expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Set Token expiration to 5 mins
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create JWT claims
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//Declare token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error creating token, return http 500
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {

}
