package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBClient struct {
	*mongo.Client
}

// ConnectToDB connects server with MongoDB, returns database client
func ConnectToDB() *mongo.Client {

	//Database config
	client, err := mongo.NewClient(options.Client()).ApplyURI("mongo_uri_here")
	if err != nil {
		log.Fatalf("[-] MongoDB NewClient error: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("[-] MongoDB client.Connect error: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatalf("[-] Ping error: %v", err)
	}

	return client
	
}

func (c *DBClient) InsertUser() {

}

// GetExpectedPassword queries password from the database based on given username
func (c *DBClient) GetExpectedPassword(username string) string {
	collection := c.Database("korero").Collection("users")
	if err = collection.FindOne()

}

// CheckIfUsernameTakes queries given username from the database and check if such entry already exists
func (c *DBClient) CheckIfUsernameTaken(username string) bool {
	collection := c.Database("korero").Collection("users")
}
