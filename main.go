package main

import (
	"log"
	"net/http"

	"github.com/korero-chat/backend/pkg/router"
)

func main() {
	router := router.SetRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
