package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config defines the global config object
type Config struct {
	MongoDB *mongo.Client
	Env     *Env
}

// Init Initializes Global Config
func Init() (*Config, error) {

	newConfig := new(Config)

	newConfig.Env = LoadEnvironment()

	newConfig.ConnectMongo()

	return newConfig, nil

}

// ConnectMongo connects mongo
func (c *Config) ConnectMongo() {

	ctx := context.Background()

	host := c.Env.MongoDBEnv.Host
	port := c.Env.MongoDBEnv.Port
	database := c.Env.MongoDBEnv.Database
	user := c.Env.MongoDBEnv.User
	password := c.Env.MongoDBEnv.Password
	srvMode := c.Env.MongoDBEnv.SRV

	mongoURI := ""

	if srvMode == "true" {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", user, password, host)
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
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
