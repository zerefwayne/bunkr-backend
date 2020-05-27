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
}

// Init Initializes Global Config
func Init() (*Config, error) {

	newConfig := new(Config)

	newConfig.ConnectMongo()

	return newConfig, nil

}

// ConnectMongo connects mongo
func (c *Config) ConnectMongo() {

	ctx := context.Background()

	host := "localhost"
	port := "27017"
	database := "collegeportal"
	user := ""
	password := ""
	srvMode := "false"

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

	log.Printf("database	| connected successfully: %s\n", mongoURI)

}
