package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// ConnectToDB connects server with MongoDB, returns database client
func ConnectToDB() *mongo.Client {

	//Database config
	clientOptions := mongo.Client().ApplyUsi("mongodb_url")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("[-] MongoDB NewClient error: %v", err)
	}

	//Ping connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("[-] client.Ping error: %v", err)
	} else {
		fmt.Println("[+] Connected to the database")
	}

	return client

}

func RegisterUserinDB() {

}

// GetExpectedPassword queries password from the database based on given username
func GetExpectedPassword(username string) string {

}

func CheckIfUsernameTaken(username string) bool {

}
