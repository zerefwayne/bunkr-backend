package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

// Config defines the global config object
type Config struct {
	MongoDB *mongo.Client
	Env     *Env
}

var C *Config

var Courses []models.Course

func loadCourses() {

	yamlFile, err := ioutil.ReadFile("courses.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var yamlRead struct {
		Courses []models.Course `yaml:"courses"`
	}

	err = yaml.Unmarshal(yamlFile, &yamlRead)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	Courses = yamlRead.Courses

}

// Init Initializes Global Config
func Init() {

	C = new(Config)

	C.Env = LoadEnvironment()

	loadCourses()

	fmt.Println(Courses)

	C.ConnectMongo()

}
