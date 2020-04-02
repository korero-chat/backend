package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	if err != nil {
		log.Fatalf("[-] Ping error: %v", err)
	}

	return client

}

// InsertUser inserts new registered user to the database
func (c *DBClient) InsertUser(username, passwordhash, email string) {
	newUser := User{
		Username:     username,
		PasswordHash: passwordhash,
		Email:        email,
	}

	collection := c.Database("korero").Collection("users")
	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatalf("[-] InsertOne User error: %v", err)
	}

}

// GetExpectedPassword queries password from the database based on given username
func (c *DBClient) GetExpectedPasswordHash(username string) string {
	var expectedPasswordHash string

	collection := c.Database("korero").Collection("users")
	result := collection.FindOne(context.Background(), bson.M{"username": username})

	_, err := result.Decode(&expectedPasswordHash)

	return expectedPasswordHash

}

// CheckIfUsernameTakes queries given username from the database and check if such entry already exists
func (c *DBClient) CheckIfUsernameTaken(username string) bool {

}
