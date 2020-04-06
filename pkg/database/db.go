package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	*mongo.Client
}

// ConnectToDB connects server with MongoDB, returns database client
func ConnectToDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("[-] Mongo.Connect error: %v", err)
	}

	fmt.Println("[+] Connected to the database")

	return client
}

// InsertUser inserts new registered user to the database
func InsertUser(username, passwordhash, email string) {
	c := ConnectToDB()

	newUser := User{
		Username:       username,
		PasswordHash:   passwordhash,
		Email:          email,
		EmailConfirmed: false,
		Chats:          nil,
	}

	collection := c.Database("korero").Collection("users")
	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatalf("[-] InsertOne User error: %v", err)
	}

}

// GetExpectedPasswordHash queries password from the database based on given username
func GetExpectedPasswordHash(username string) string {
	c := ConnectToDB()

	var expectedPasswordHash string

	collection := c.Database("korero").Collection("users")
	result := collection.FindOne(context.TODO(), bson.M{"username": username})

	err := result.Decode(&expectedPasswordHash)
	if err != nil {
		log.Fatalf("[-] GetExpectedPasswordHash Decode error: %v", err)
	}

	return expectedPasswordHash

}

// CheckIfUsernameTaken queries given username from the database and check if such entry already exists
func CheckIfEmailTaken(email string) bool {
	c := ConnectToDB()

	var user User

	collection := c.Database("korero").Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{"username": email}).Decode(&user)
	if err != nil {
		return true
	}

	if email == user.Email {
		return false
	}

	return true
}
