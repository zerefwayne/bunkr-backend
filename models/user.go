package models

// User Defines the user model
type User struct {
	ID                string   `json:"id" bson:"_id"`
	Username          string   `json:"username" bson:"username"`
	Name              string   `json:"name" bson:"name"`
	Password          string   `json:"password,omitempty" bson:"password"`
	Email             string   `json:"email" bson:"email"`
	IsAdmin           bool     `json:"isAdmin" bson:"isAdmin"`
	SubscribedCourses []string `json:"subscribedCourses" bson:"subscribedCourses"`
	Bookmarks         []string `json:"bookmarks" bson:"bookmarks"`
}
