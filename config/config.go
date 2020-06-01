package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Config defines the global config object
type Config struct {
	MongoDB *mongo.Client
	Env     *Env
}

var C *Config

// Init Initializes Global Config
func Init() {

	C = new(Config)

	C.Env = LoadEnvironment()

	C.ConnectMongo()

}
