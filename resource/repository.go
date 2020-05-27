package resource

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Repository defines repository interface for Resource
type Repository interface {
	CreateResource(ctx context.Context, resource *models.Resource) error
	GetResourcesAll(ctx context.Context) ([]*models.Resource, error)
	GetResourceByID(ctx context.Context, id string) (*models.Resource, error)
	GetResourceByUserID(ctx context.Context, id string) ([]*models.Resource, error)
}

type resourceRepository struct {
}

func NewResourceRepository() Repository {
	return &resourceRepository{}
}

func (r *resourceRepository) CreateResource(ctx context.Context, resource *models.Resource) error {
	return nil
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
