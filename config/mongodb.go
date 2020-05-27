package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo connects mongo
func (c *Config) ConnectMongo() {

	ctx := context.Background()

	host := c.Env.MongoDBEnv.Host
	port := c.Env.MongoDBEnv.Port
	user := c.Env.MongoDBEnv.User
	password := c.Env.MongoDBEnv.Password
	srvMode := c.Env.MongoDBEnv.SRV

	mongoURI := ""

	if srvMode == "true" {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", user, password, host)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoURI,
	))

	if err != nil {
		log.Fatal(err)
	}

	c.MongoDB = client

	if err := c.MongoDB.Ping(context.Background(), nil); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("database	| connected successfully: %s\n", mongoURI)
	}
}
