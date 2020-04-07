package router

import (
	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/pkg/api"
	"github.com/korero-chat/backend/pkg/auth"
)

// SetRouter Creates mux Router instance to handle api routes
func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", api.RegisterUserEndpoint).Methods("POST")
	router.HandleFunc("/api/chats", api.GetUsersChatsEndpoint).Methods("GET")
	router.HandleFunc("/api/chat", api.CreateNewChatRoomEndpoint).Methods("POST")
	//router.HandleFunc("/api/chat/{id}, api.EnterChatRoomEndpoint").Methods("GET")
	router.HandleFunc("/api/chat/{id}/users", api.GetUsersOfChatEndpoint).Methods("GET")

	// auth
	router.HandleFunc("/signin", auth.SignIn)
	router.HandleFunc("/refresh", auth.Refresh)
	router.HandleFunc("/welcome", auth.Welcome)
	return router
}
