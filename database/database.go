package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/korero-chat/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURI, dbName string

func init() {
	err := godotenv.Load()
	if err == nil {
		// .env exists
		log.Println(".env file has been loaded successfully.")
		mongoURI = os.Getenv("MONGO_URI")
		dbName = os.Getenv("DBNAME")
	} else {
		// .env does not exist
		log.Println("Note: Cannot load .env file. Using test credentials.")
		mongoURI = "mongodb://localhost/"
		dbName = "test"
	}
}

func ConnectToDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("[-] Mongo.Connect error: %v", err)
	}

	fmt.Println("[+] Connected to the database")

	return client
}

func FindUserByUsername(username string) (models.User, error) {
	c := ConnectToDB()

	var user models.User

	collection := c.Database(dbName).Collection("users")
	err := collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUser(user models.User) error {
	c := ConnectToDB()
	collection := c.Database(dbName).Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)

	return err

}

func InsertChat(chat models.Chat) error {
	c := ConnectToDB()
	collection := c.Database(dbName).Collection("chats")
	_, err := collection.InsertOne(context.TODO(), chat)

	return err
}

func GetChatByID(chatID string) (models.Chat, error) {
	c := ConnectToDB()
	collection := c.Database(dbName).Collection("chats")

	var chat models.Chat
	err := collection.FindOne(context.TODO(), bson.M{"id": chatID}).Decode(&chat)
	if err != nil {
		log.Fatalf("[-] Error while querying chat: %v", err)
	}

	return chat, err
}
