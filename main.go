package main

import (
	"github.com/zerefwayne/college-portal-backend/api"
	"github.com/zerefwayne/college-portal-backend/config"
)

func init() {
	config.Init()
}

func main() {
	api.Serve()
}
