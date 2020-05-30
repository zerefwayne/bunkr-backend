package models

type Course struct {
	Slug      string   `json:"slug,omitempty" yaml:"slug"`
	Code      string   `json:"code,omitempty" yaml:"code"`
	Name      string   `json:"name,omitempty" yaml:"name"`
	Resources []string `json:"resources" yaml:"resources"`
}
