package resource

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type resourceRepository struct {
	DB *mongo.Database
}

// NewMongoResourceRepository ...
func NewMongoResourceRepository(client *mongo.Client) Repository {
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
	return nil, nil
}

func (r *resourceRepository) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	return nil, nil
}

func (r *resourceRepository) GetResourceByUserID(ctx context.Context, id string) ([]*models.Resource, error) {
	return nil, nil
}
