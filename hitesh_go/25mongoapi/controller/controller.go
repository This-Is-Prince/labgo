package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("No Error")
	}

	// client option
	clientOption := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection reference / instance
	fmt.Println("Collection reference / instance is ready")
}
