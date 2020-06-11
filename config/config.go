package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Config defines the global config object
type Config struct {
	MongoDB *mongo.Client
	Env     *Env
}

// C Global Config object
var C *Config

// Init Initializes Global Config
func Init() {

	// Creates a new config object
	C = new(Config)

	// Loads environment variables
	C.Env = LoadEnvironment()

	// MongoDB connection is created
	C.ConnectMongo()

}
