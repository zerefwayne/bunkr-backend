package main

import (
	"log"

	"github.com/zerefwayne/college-portal-backend/config"
)

// Config defines the global config
var Config *config.Config

func init() {

	Config, err := config.Init()

	if err != nil {
		log.Fatalln(err)
	}

	_ = Config

}

func main() {
	select {}
}
