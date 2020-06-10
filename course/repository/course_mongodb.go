package repository

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type courseRepository struct {
	DB *mongo.Database
}

// NewMongoResourceRepository ...
func NewMongoResourceRepository(client *mongo.Client) course.Repository {
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

func (r *courseRepository) GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error) {

	var course models.Course

	filter := bson.M{"slug": slug}

	collection := r.DB.Collection("courses")

	err := collection.FindOne(ctx, filter).Decode(&course)

	return &course, err

}

func (r *courseRepository) CreateCourse(ctx context.Context, course *models.Course) error {

	collection := r.DB.Collection("courses")

	_, err := collection.InsertOne(ctx, course)

	return err

}

func (r *courseRepository) PushResource(ctx context.Context, courseCode string, resourceID string) error {

	collection := r.DB.Collection("courses")

	filter := bson.M{"code": courseCode}

	update := bson.M{"$addToSet": bson.M{"resource_ids": resourceID}}

	result := collection.FindOneAndUpdate(ctx, filter, update)

	return result.Err()

}

func (r *courseRepository) PopResource(ctx context.Context, resourceID string) error {

	collection := r.DB.Collection("courses")

	filter := bson.M{}

	update := bson.M{"$pull": bson.M{"resource_ids": resourceID}}

	_, err := collection.UpdateMany(ctx, filter, update)

	return err

}
