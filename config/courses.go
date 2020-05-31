package config

import (
	"io/ioutil"
	"log"

	"github.com/zerefwayne/college-portal-backend/models"
	"gopkg.in/yaml.v2"
)

var Courses []*models.Course

func loadCourses() {

	yamlFile, err := ioutil.ReadFile("courses.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var yamlRead struct {
		Courses []*models.Course `yaml:"courses"`
	}

	err = yaml.Unmarshal(yamlFile, &yamlRead)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	Courses = yamlRead.Courses

}
