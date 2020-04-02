package router

import (
	"github.com/gorilla/mux"
	"github.com/korero-chat/backend/pkg/api"
)

// SetRouter Creates mux Router instance to handle api routes
func SetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/newuser", api.RegisterUserEndpoint).Methods("POST")

	return router
}
