package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nickWoott/my-hex-api/internal/core/domain"
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

func (db *MongoDbDatabase) GetStoryPointById(id string) (domain.StorypointRequest, error) {

	coll := db.client.Database("storypoint-store").Collection("storypoints")

	var result domain.StorypointRequest

	err := coll.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found id of: %s\n", id)
		return result, err
	}

	if err != nil {

		panic(err)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {

		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return result, nil
}


//does this also use domain types?
//I think you can use the domain types in any place, domain types will not change
func (db *MongoDbDatabase) PostStoryPoint(storypoint domain.StorypointRequest) error {

 coll := db.client.Database("storypoint-store").Collection("storypoints")
    
    // Convert domain.StoryPoint to BSON document
    document := bson.M{
        "id":      storypoint.Id,
        "text":    storypoint.Text,
        "choices": storypoint.Choices,
        // Add any other fields from domain.StoryPoint as needed
    }
    
    _, err := coll.InsertOne(context.TODO(), document)

    if err != nil {
        fmt.Printf("Failed to insert story point into database: %v\n", err)
        return err
    }
    
    return nil
}