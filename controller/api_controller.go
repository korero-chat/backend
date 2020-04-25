package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/database"
	"github.com/korero-chat/backend/models"
	"gopkg.in/validator.v2"
)

func CreateChatEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var chat models.Chat
	var response models.ResponseModel

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &chat)
	if err != nil {
		w.WriteHeader(500)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	//verify chat name
	nur := models.NewChatRequest{Name: chat.Name}
	if errs := validator.Validate(nur); errs != nil {
		w.WriteHeader(422)
		response.Error = errs.Error()
		response.Result = "Validation error"
		json.NewEncoder(w).Encode(response)
		return
	}

	//insert chat into database
	err = database.InsertChat(chat)
	if err != nil {
		w.WriteHeader(500)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(200)
	response.Result = "Chat created successfully"
	json.NewEncoder(w).Encode(response)
}

func GetChatEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response models.ResponseModel

	params := mux.Vars(r)
	chat, err := database.GetChatByID(params["id"])
	if err != nil {
		w.WriteHeader(500)
		response.Error = err.Error()
		json.NewEncoder(w).Encode(response)
		return

	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(chat)
	return
}
