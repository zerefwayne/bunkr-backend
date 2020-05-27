package resource

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Usecase defines usecase interface for Resource
type Usecase interface {
	CreateResource(ctx context.Context, resource *models.Resource) error
	GetResourcesAll(ctx context.Context) ([]*models.Resource, error)
	GetResourceByID(ctx context.Context, id string) (*models.Resource, error)
	GetResourceByUserID(ctx context.Context, id string) ([]*models.Resource, error)
}
