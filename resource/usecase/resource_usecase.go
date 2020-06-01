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

type resourceUsecase struct {
	resourceRepo resource.Repository
	course       course.Usecase
}

func NewResourceUsecase() resource.Usecase {

	return &resourceUsecase{
		resourceRepo: repository.NewMongoResourceRepository(config.C.MongoDB),
		course:       common.Course,
	}

}

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
	resource, err := u.resourceRepo.GetResourceByID(ctx, id)
	return resource, err
}

func (u *resourceUsecase) GetResourcesByUserID(ctx context.Context, id string) ([]*models.Resource, error) {

	resources, err := u.resourceRepo.GetResourcesByUserID(ctx, id)

	return resources, err
}

func (u *resourceUsecase) DeleteResourceByID(ctx context.Context, id string) error {

	err := u.resourceRepo.DeleteResourceByID(ctx, id)

	return err

}
