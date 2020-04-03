package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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

	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[-] error loading .env file: %v", err)
	}

	//Database config
	clientoptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientoptions)
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
		log.Fatalf("[-] CheckIfEmailTaken error: %v", err)
	}

	if email == user.Email {
		return false
	}

	return true
}
