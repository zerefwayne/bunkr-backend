package resource

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type resourceRepository struct {
	DB *mongo.Database
}

// NewMongoResourceRepository ...
func newMongoResourceRepository(client *mongo.Client) Repository {
	database := client.Database(config.C.Env.Database)

	return &resourceRepository{
		DB: database,
	}
}

func (r *resourceRepository) CreateResource(ctx context.Context, resource *models.Resource) error {

	collection := r.DB.Collection("resources")

	_, err := collection.InsertOne(ctx, resource)

	return err
}

func (r *resourceRepository) GetResourcesAll(ctx context.Context) ([]*models.Resource, error) {
	collection := r.DB.Collection("resources")

	filter := bson.M{}

	results, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var resources []*models.Resource

	for results.Next(ctx) {

		resource := new(models.Resource)

		if err := results.Decode(resource); err != nil {
			log.Println(err)
			return resources, err
		}

		resources = append(resources, resource)

	}

	return resources, nil
}

func (r *resourceRepository) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	return nil, nil
}

func (r *resourceRepository) GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error) {

	collection := r.DB.Collection("resources")

	filter := bson.M{"created_by": id}

	results, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var resources []*models.Resource

	for results.Next(ctx) {

		resource := new(models.Resource)

		if err := results.Decode(resource); err != nil {
			log.Println(err)
			return resources, err
		}

		resources = append(resources, resource)

	}

	return resources, nil
}

func (r *resourceRepository) DeleteResourceByID(ctx context.Context, id string) error {

	collection := r.DB.Collection("resources")

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return errors.New("document doesn't exist")
	}

	return err
}
