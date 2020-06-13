package config

import (
	"github.com/sendgrid/sendgrid-go"
	"go.mongodb.org/mongo-driver/mongo"
)

// Config defines the global config object
type Config struct {
	MongoDB  *mongo.Client
	Env      *Env
	SendGrid *sendgrid.Client
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

	// SendGrid connection
	C.ConnectSendGrid()

}
