package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/zerefwayne/college-portal-backend/common"
	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/course/repository"
	"github.com/zerefwayne/college-portal-backend/models"
)

type CourseUsecase struct {
	CourseRepo course.Repository
}

func NewCourseUsecase() *CourseUsecase {
	return &CourseUsecase{
		CourseRepo: repository.NewMongoResourceRepository(config.C.MongoDB),
	}
}

func (u *CourseUsecase) PushResource(ctx context.Context, courseCode string, resourceID string) error {

	course, err := u.GetCourseByCode(ctx, courseCode)

	if err != nil {
		return err
	}

	fmt.Println(course)

	if err := u.CourseRepo.PushResource(ctx, course.Code, resourceID); err != nil {
		return err
	}

	return nil

}

func (u *CourseUsecase) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	courses, err := u.CourseRepo.GetAllCourses(ctx)
	return courses, err
}

func (u *CourseUsecase) CreateCourse(ctx context.Context, course *models.Course) error {

	err := u.CourseRepo.CreateCourse(ctx, course)

	return err

}

func (u *CourseUsecase) GetCourseByCode(ctx context.Context, code string) (*models.Course, error) {

	course, err := u.CourseRepo.GetCourseByCode(ctx, code)

	if err != nil {
		return nil, errors.New("course not found")
	}

	for _, resourceID := range course.ResourceIDs {

		log.Println(resourceID)

		resource, err := common.Resource.GetResourceByID(ctx, resourceID)

		if err != nil {
			return nil, err
		}

		course.Resources = append(course.Resources, resource)

	}

	return course, nil
}

func (u *CourseUsecase) GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error) {

	course, err := u.CourseRepo.GetCourseBySlug(ctx, slug)

	if err != nil {
		return nil, errors.New("course not found")
	}

	for _, resourceID := range course.ResourceIDs {

		log.Println(resourceID)

		resource, err := common.Resource.GetResourceByID(ctx, resourceID)

		if err != nil {
			return nil, err
		}

		course.Resources = append(course.Resources, resource)

	}

	return course, nil
}
