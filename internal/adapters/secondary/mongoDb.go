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

func (db *MongoDbDatabase) GetStoryPointById(id int) (domain.StorypointRequest, error) {

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
func (db *MongoDbDatabase) PostStoryPoints(storyPoints []domain.StorypointRequest) error {

 coll := db.client.Database("storypoint-store").Collection("storypoints")
 
 for _, sp := range storyPoints {
	filter := bson.M{"_id": sp.Id} // Assuming Id is the unique identifier

	// Create an update operation to set the document
	update := bson.M{"$set": sp}

	// Perform the update operation with upsert set to true
	_, err := coll.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Printf("Failed to update story point in database: %v\n", err)
		return err
	}
}

return nil
}

func(db *MongoDbDatabase) DeleteAllStoryPoints () error {

	coll := db.client.Database("storypoint-store").Collection("storypoints")

	_, err := coll.DeleteMany(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println("there was an error deleting all documents")
	}
	fmt.Println("all documents deleted from the collection")

	return nil
}