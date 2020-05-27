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

type resourceUsecase struct {
	resourceRepo Repository
}

// NewResourceUsecase ...
func NewResourceUsecase(r Repository) Repository {
	return &resourceUsecase{
		resourceRepo: r,
	}
}

func (u *resourceUsecase) CreateResource(ctx context.Context, resource *models.Resource) error {
	return nil
}

func (u *resourceUsecase) GetResourcesAll(ctx context.Context) ([]*models.Resource, error) {
	return nil, nil
}

func (u *resourceUsecase) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	return nil, nil
}

func (u *resourceUsecase) GetResourceByUserID(ctx context.Context, id string) ([]*models.Resource, error) {
	return nil, nil
}
