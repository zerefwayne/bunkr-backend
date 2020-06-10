package resource

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Usecase defines usecase interface for Resource
type Usecase interface {
	CreateResource(ctx context.Context, resource *models.Resource) error
	UpdateResource(ctx context.Context, resource *models.Resource) error
	GetResourcesAll(ctx context.Context) ([]*models.Resource, error)
	GetPendingResources(ctx context.Context) ([]*models.Resource, error)
	GetResourceByID(ctx context.Context, id string) (*models.Resource, error)
	GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error)
	DeleteResourceByID(ctx context.Context, id string) error
	ApproveResourceByID(ctx context.Context, id string) error
	AddVoteResource(ctx context.Context, resourceID string, userID string) error
	UpdateVoteResource(ctx context.Context, resourceID string, userID string) error
}
