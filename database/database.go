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

func ConnectToDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName = os.Getenv("DBNAME")
	mongoURI = os.Getenv("MONGO_URI")

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
