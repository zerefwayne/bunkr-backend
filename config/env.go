package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// MongoDBEnv Loads MongoDB Environment Variables
type MongoDBEnv struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SRV      string
}

// Load Loads MongoDB Environment Variables
func (me *MongoDBEnv) Load() {
	me.Host = os.Getenv("MONGODB_HOST")
	me.Port = os.Getenv("MONGODB_PORT")
	me.Database = os.Getenv("MONGODB_DATABASE")
	me.User = os.Getenv("MONGODB_USER")
	me.Password = os.Getenv("MONGODB_PASSWORD")
	me.SRV = os.Getenv("MONGODB_SRV")
}

type APIEnv struct {
	Port string
}

func (a *APIEnv) Load() {
	a.Port = os.Getenv("API_PORT")
}

// Env ...
type Env struct {
	MongoDBEnv
	APIEnv
}

// LoadEnvironment Loads environment variables
func LoadEnvironment() *Env {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
		return nil
	}

	env := new(Env)

	env.APIEnv.Load()
	env.MongoDBEnv.Load()

	return env

}
