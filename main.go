package main

import (
	"github.com/korero-chat/backend/pkg/database"
)

func main() {
	client := database.ConnectToDB()
}
