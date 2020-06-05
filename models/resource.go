package models

import "time"

// Resource Defines the Resource Schema
type Resource struct {
	ID         string    `json:"id" bson:"_id"`
	Type       string    `json:"type" bson:"type"`
	Content    string    `json:"content" bson:"content"`
	Title      string    `json:"title" bson:"title"`
	Tags       []string  `json:"tags" bson:"tags"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	CreatedBy  string    `json:"created_by" bson:"created_by"`
	IsApproved bool      `json:"is_approved" bson:"is_approved"`
}

// Allowed values for type: link, article, file
