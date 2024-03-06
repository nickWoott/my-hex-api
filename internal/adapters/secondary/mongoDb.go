package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//define the struct



type MongoDbDatabase struct {
	client *mongo.Client
}

// write a function that instantiates the struct
func NewMongoDbDatabase() MongoDbDatabase {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		println("error connecting")
		panic(err)
	}

	return MongoDbDatabase{
		client: client,
	}

}

func (db *MongoDbDatabase) GetStoryPointById(id string) string {

	coll := db.client.Database("storypoint-store").Collection("storpoints")

	var result bson.M

	err := coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found id of: %s\n", id)
		return "function did not work"
	}

	if err != nil {

		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {

		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return string(jsonData)
}
