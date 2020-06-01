package models

type Course struct {
	Slug        string      `json:"slug,omitempty" yaml:"slug" bson:"slug"`
	Code        string      `json:"code,omitempty" yaml:"code" bson:"code"`
	Name        string      `json:"name,omitempty" yaml:"name" bson:"name"`
	ResourceIDs []string    `json:"resource_ids" yaml:"resource_ids" bson:"resource_ids"`
	Resources   []*Resource `json:"resources" bson:"resources"`
}
