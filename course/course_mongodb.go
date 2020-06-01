package course

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type courseRepository struct {
	DB *mongo.Database
}

// NewMongoResourceRepository ...
func newMongoResourceRepository(client *mongo.Client) Repository {
	database := client.Database(config.C.Env.Database)

	return &courseRepository{
		DB: database,
	}
}

func (r *courseRepository) GetAllCourses(ctx context.Context) ([]*models.Course, error) {

	var courses []*models.Course

	filter := bson.M{}

	collection := r.DB.Collection("courses")

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return courses, err
	}

	for cursor.Next(ctx) {

		var course models.Course

		if err := cursor.Decode(&course); err != nil {
			return courses, err
		}

		courses = append(courses, &course)

	}

	return courses, nil
}

func (r *courseRepository) GetCourseByCode(ctx context.Context, code string) (*models.Course, error) {

	var course models.Course

	filter := bson.M{"code": code}

	collection := r.DB.Collection("courses")

	err := collection.FindOne(ctx, filter).Decode(&course)

	return &course, err
}

func (r *courseRepository) CreateCourse(ctx context.Context, course *models.Course) error {

	collection := r.DB.Collection("courses")

	_, err := collection.InsertOne(ctx, course)

	return err

}
