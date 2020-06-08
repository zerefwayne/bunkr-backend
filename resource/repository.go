package resource

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Repository defines repository interface for Resource
type Repository interface {
	CreateResource(ctx context.Context, resource *models.Resource) error
	GetResourcesAll(ctx context.Context) ([]*models.Resource, error)
	GetPendingResources(ctx context.Context) ([]*models.Resource, error)
	GetResourceByID(ctx context.Context, id string) (*models.Resource, error)
	GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error)
	DeleteResourceByID(ctx context.Context, id string) error
	ApproveResourceByID(ctx context.Context, id string) error
	UpdateResource(ctx context.Context, resource *models.Resource) error
}
