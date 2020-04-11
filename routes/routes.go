package routes

import (
	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/controller"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	//Auth routes
	router.HandleFunc("/register", controller.RegisterUserEndpoint).Methods("POST")
	router.HandleFunc("/login", controller.LoginEndpoint).Methods("POST")
	//router.HandleFunc("/logout", controller.LogoutEndpoint)

	return router
}
