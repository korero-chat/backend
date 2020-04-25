package routes

import (
	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/controller"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	//Auth routes
	router.HandleFunc("/backend/register", controller.RegisterUserEndpoint).Methods("POST")
	router.HandleFunc("/backend/login", controller.LoginEndpoint).Methods("POST")

	// subroutes
	s := router.PathPrefix("/api").Subrouter()
	s.Use(controller.VerifyToken)
	s.HandleFunc("/api/chat", controller.CreateChatEndpoint).Methods("POST")
	s.HandleFunc("/api/chat/{id}", controller.GetChatEndpoint).Methods("GET")
	return router
}
