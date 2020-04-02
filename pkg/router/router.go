package router

import (
	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/pkg/api"
)

// SetRouter Creates mux Router instance to handle api routes
func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/newuser", api.RegisterUserEndpoint).Methods("POST")
	router.HandleFunc("/api/user/chat/{id}", api.GetChatsByUserIDEndpoint).Methods("GET")
	router.HandleFunc("/api/chat/{id}/users", api.GetChatUsersEndpoint).Methods("GET")

	return router
}
