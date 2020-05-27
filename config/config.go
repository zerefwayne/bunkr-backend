package config

import (
	"go.mongodb.org/mongo-driver/mongo"
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
