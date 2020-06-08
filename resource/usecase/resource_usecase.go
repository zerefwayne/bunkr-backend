package usecase

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/zerefwayne/college-portal-backend/common"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource"
	"github.com/zerefwayne/college-portal-backend/resource/repository"
)

type ResourceUsecase struct {
	ResourceRepo resource.Repository
	Course       course.Usecase
}

func NewResourceUsecase() *ResourceUsecase {

	return &ResourceUsecase{
		ResourceRepo: repository.NewMongoResourceRepository(config.C.MongoDB),
		Course:       common.Course,
	}

}

func (u *ResourceUsecase) CreateResource(ctx context.Context, resource *models.Resource) error {

	resource.ID = xid.New().String()
	resource.CreatedAt = time.Now()

	err := u.ResourceRepo.CreateResource(ctx, resource)

	return err
}

func (u *ResourceUsecase) UpdateResource(ctx context.Context, resource *models.Resource) error {
	err := u.ResourceRepo.UpdateResource(ctx, resource)
	return err
}

func (u *ResourceUsecase) GetResourcesAll(ctx context.Context) ([]*models.Resource, error) {
	resources, err := u.ResourceRepo.GetResourcesAll(ctx)

	return resources, err
}

func (u *ResourceUsecase) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	resource, err := u.ResourceRepo.GetResourceByID(ctx, id)
	return resource, err
}

func (u *ResourceUsecase) GetPendingResources(ctx context.Context) ([]*models.Resource, error) {
	resource, err := u.ResourceRepo.GetPendingResources(ctx)
	return resource, err
}

func (u *ResourceUsecase) GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error) {

	resources, err := u.ResourceRepo.GetResourcesByUserID(ctx, id)

	return resources, err
}

func (u *ResourceUsecase) DeleteResourceByID(ctx context.Context, id string) error {

	err := u.ResourceRepo.DeleteResourceByID(ctx, id)

	return err

}

func (u *ResourceUsecase) ApproveResourceByID(ctx context.Context, id string) error {

	err := u.ResourceRepo.ApproveResourceByID(ctx, id)

	return err

}
