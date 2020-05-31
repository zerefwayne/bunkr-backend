package resource

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/zerefwayne/college-portal-backend/models"
)

// Usecase defines usecase interface for Resource
type Usecase interface {
	CreateResource(ctx context.Context, resource *models.Resource) error
	GetResourcesAll(ctx context.Context) ([]*models.Resource, error)
	GetResourceByID(ctx context.Context, id string) (*models.Resource, error)
	GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error)
	DeleteResourceByID(ctx context.Context, id string) error
}

type resourceUsecase struct {
	resourceRepo Repository
}

var ResourceUsecase resourceUsecase

func (u *resourceUsecase) CreateResource(ctx context.Context, resource *models.Resource) error {

	resource.ID = xid.New().String()
	resource.CreatedAt = time.Now()

	err := u.resourceRepo.CreateResource(ctx, resource)

	return err
}

func (u *resourceUsecase) GetResourcesAll(ctx context.Context) ([]*models.Resource, error) {
	resources, err := u.resourceRepo.GetResourcesAll(ctx)

	return resources, err
}

func (u *resourceUsecase) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	return nil, nil
}

func (u *resourceUsecase) GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error) {

	resources, err := u.resourceRepo.GetResourcesByUserID(ctx, id)

	return resources, err
}

func (u *resourceUsecase) DeleteResourceByID(ctx context.Context, id string) error {

	err := u.resourceRepo.DeleteResourceByID(ctx, id)

	return err

}
