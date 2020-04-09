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

func ConnectToDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal{"Error loading .env file"}
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

func FindUserByUsername(username string) error {
	c := ConnectToDB()

	collection := c.Database("korero").Collection("users")
	err := collection.FindOne(context.TODO(), bson.D{{"username", username}})

	return err
}

func InsertUser(user models.User) error {
	c := ConnectToDB()
	collection := c.Database("korero").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)

	return err

}
