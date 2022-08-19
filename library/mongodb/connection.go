package mongodb

import (
	"context"
	"fmt"
	"postinger/config"
	"postinger/util/logwrapper"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection - Connection structure
type Connection struct {
	Conn     *mongo.Client
	ConnDB   *mongo.Database
	Database string
}

// Client - MongoDB Connection
var Client *Connection

// NewConnection - new connection of amqp
func NewConnection(mongoConfig config.MongoDBConfig) error {
	if mongoConfig.URL == "" || mongoConfig.Database == "" {
		return fmt.Errorf("CONFIGURATION IS MISSING FOR MONGODB")
	}

	mongoClient := &Connection{
		Conn:     nil,
		ConnDB:   nil,
		Database: mongoConfig.Database,
	}
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoConfig.URL)

	var err error
	mongoClient.Conn, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("MONGO ERROR: %s", err)
	}
	err = mongoClient.Conn.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("MONGO PING ERROR: %s", err)
	}

	logwrapper.Logger.Infoln("Connected to MONGODB_URL : ", mongoConfig.URL)

	mongoClient.ConnDB = mongoClient.Conn.Database(mongoClient.Database)
	Client = mongoClient

	return nil
}

// GetCollection - Helper Functions
func GetCollection(collectionName string) *mongo.Collection {
	return Client.ConnDB.Collection(collectionName)
}

// DbContext - Helper Functions
func DbContext(i time.Duration) (context.Context, context.CancelFunc) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), i*time.Second)
	return ctx, ctxCancel
}
