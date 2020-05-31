package config

import (
	"fmt"

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

	loadCourses()

	fmt.Println(Courses)

	C.ConnectMongo()

}
