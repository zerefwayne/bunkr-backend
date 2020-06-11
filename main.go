package main

import (
	"github.com/zerefwayne/college-portal-backend/api"
	"github.com/zerefwayne/college-portal-backend/config"
)

func init() {
	// Initializes the config package and defines the global configuration
	config.Init()
}

func main() {
	// Starts the REST API
	api.Serve()
}
