package mongoService

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConnection struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect() MongoConnection {
	mongoConnection := MongoConnection{}
	mongoConnection.setClient()
	mongoConnection.setDatabase()

	return mongoConnection
}

func ContextTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (c *MongoConnection) setClient() {
	clientOptions := options.Client().ApplyURI("mongodb://mongoadmin:secret@localhost:27017/?authSource=admin")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Mongo Connected")
	c.Client = client
}

func (c *MongoConnection) setDatabase() {
	c.Database = c.Client.Database("pokedex")
}
