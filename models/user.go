package models

// User Defines the user model
type User struct {
	ID       string `json:"_id", bson:"_id"`
	Username string `json:"username", bson:"username"`
	Name     string `json:"name", bson:"name"`
	Password string `json:"password", bson:"password"`
	Email    string `json:"email", bson:"email"`
}
