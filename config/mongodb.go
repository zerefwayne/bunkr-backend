package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Returns the MongoURI from environment
func (c *Config) generateMongoURI() string {

	host := c.Env.MongoDBEnv.Host
	port := c.Env.MongoDBEnv.Port
	database := c.Env.MongoDBEnv.Database
	user := c.Env.MongoDBEnv.User
	password := c.Env.MongoDBEnv.Password
	srvMode := c.Env.MongoDBEnv.SRV

	mongoURI := ""

	if srvMode == "true" {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@cluster0-bevsh.mongodb.net/%s?retryWrites=true&w=majority", user, password, database)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	return mongoURI
}

// ConnectMongo connects mongo
func (c *Config) ConnectMongo() {

	ctx := context.Background()

	mongoURI := c.generateMongoURI()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))

	if err != nil {
		log.Fatal(err)
	}

	c.MongoDB = client

	// Pings the client to ensure a proper connection
	if err := c.MongoDB.Ping(context.Background(), nil); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("database	| connected successfully: %s\n", mongoURI)
	}
}
