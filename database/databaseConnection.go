package database

import(
	"fmt"
	"time"
	"log"
	"os"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongoClient{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDb := os.Getenv("MONGODB_DB")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err !=nil {
		log.Fatal(err)
	}

	context.WithTimeout(context.Background(), 10*time.Second) 
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongo")

	return Client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongodb.Client, collectionName string) *mongoCollection{
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectName)
	return collection
}