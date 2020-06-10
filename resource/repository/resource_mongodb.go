package repository

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type resourceRepository struct {
	DB *mongo.Database
}

// NewMongoResourceRepository ...
func NewMongoResourceRepository(client *mongo.Client) resource.Repository {
	database := client.Database(config.C.Env.Database)

	return &resourceRepository{
		DB: database,
	}
}

func (r *resourceRepository) AddVoteResource(ctx context.Context, resourceID string, userID string) error {

	collection := r.DB.Collection("resources")

	filter := bson.M{"_id": resourceID}

	update := bson.M{"$addToSet": bson.M{"upvotes": userID}}

	result := collection.FindOneAndUpdate(ctx, filter, update)

	return result.Err()

}

func (r *resourceRepository) UpdateVoteResource(ctx context.Context, resourceID string, userID string) error {

	collection := r.DB.Collection("resources")

	filter := bson.M{"_id": resourceID}

	update := bson.M{"$pull": bson.M{"upvotes": userID}}

	result := collection.FindOneAndUpdate(ctx, filter, update)

	return result.Err()

}

func (r *resourceRepository) CreateResource(ctx context.Context, resource *models.Resource) error {

	collection := r.DB.Collection("resources")

	_, err := collection.InsertOne(ctx, resource)

	return err
}

func (r *resourceRepository) UpdateResource(ctx context.Context, resource *models.Resource) error {

	filter := bson.M{"_id": resource.ID}

	update := bson.M{"$set": bson.M{"title": resource.Title, "type": resource.Type, "content": resource.Content}}

	collection := r.DB.Collection("resources")

	err := collection.FindOneAndUpdate(ctx, filter, update)

	return err.Err()
}

func (r *resourceRepository) GetPendingResources(ctx context.Context) ([]*models.Resource, error) {
	collection := r.DB.Collection("resources")

	filter := bson.M{"is_approved": false}

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

	collection := r.DB.Collection("resources")

	filter := bson.M{"_id": id}

	var resource models.Resource

	err := collection.FindOne(ctx, filter).Decode(&resource)

	if err != nil {

		return nil, err
	}

	if len(resource.Upvotes) == 0 {
		resource.Upvotes = []string{}
	}

	return &resource, nil
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

		var resource models.Resource

		if err := results.Decode(&resource); err != nil {
			return resources, err
		}

		resources = append(resources, &resource)

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

func (r *resourceRepository) ApproveResourceByID(ctx context.Context, id string) error {

	collection := r.DB.Collection("resources")

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"is_approved": true}}

	updateResult := collection.FindOneAndUpdate(ctx, filter, update)

	return updateResult.Err()
}
