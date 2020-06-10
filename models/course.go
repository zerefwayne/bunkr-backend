package models

type Course struct {
	Slug        string      `json:"slug,omitempty" bson:"slug"`
	Code        string      `json:"code,omitempty" bson:"code"`
	Name        string      `json:"name,omitempty" bson:"name"`
	ResourceIDs []string    `json:"resource_ids" bson:"resource_ids"`
	Resources   []*Resource `json:"resources" bson:"-"`
}
