package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// MongoDBEnv Loads MongoDB Environment Variables
type MongoDBEnv struct {
	URL      string
	Host     string
	Port     string
	Database string
	User     string
	Password string
	SRV      string
}

// Load Loads MongoDB Environment Variables
func (me *MongoDBEnv) load() {
	me.Host = os.Getenv("MONGODB_HOST")
	me.Port = os.Getenv("MONGODB_PORT")
	me.Database = os.Getenv("MONGODB_DATABASE")
	me.User = os.Getenv("MONGODB_USER")
	me.Password = os.Getenv("MONGODB_PASSWORD")
	me.SRV = os.Getenv("MONGODB_SRV")
	me.URL = os.Getenv("MONGODB_URL")
}

type APIEnv struct {
	APIPort    string
	Port       string
	SigningKey string
}

func (a *APIEnv) load() {
	a.APIPort = os.Getenv("API_PORT")
	a.Port = os.Getenv("PORT")
	a.SigningKey = os.Getenv("API_SIGNING_KEY")
}

type SendGridEnv struct {
	Key            string
	BunkrClientURL string
}

func (s *SendGridEnv) load() {
	s.Key = os.Getenv("SENDGRID_KEY")
	s.BunkrClientURL = os.Getenv("BUNKR_CLIENT_URL")
}

// Env joins multiple Env objects under one struct
type Env struct {
	MongoDBEnv
	APIEnv
	SendGridEnv
}

// LoadEnvironment Loads environment variables
func LoadEnvironment() *Env {

	// Load() function looks for .env in project root by default
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
		return nil
	}

	// env is initalised
	env := new(Env)

	env.APIEnv.load()
	env.MongoDBEnv.load()
	env.SendGridEnv.load()

	return env

}
