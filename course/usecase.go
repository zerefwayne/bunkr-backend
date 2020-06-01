package course

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource"
)

type Usecase interface {
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCourseByCode(ctx context.Context, code string) (*models.Course, error)
	CreateCourse(ctx context.Context, course *models.Course) error
}

type courseUsecase struct {
	courseRepo Repository
}

var CourseUsecase courseUsecase

func (u *courseUsecase) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	courses, err := u.courseRepo.GetAllCourses(ctx)
	return courses, err
}

func (u *courseUsecase) CreateCourse(ctx context.Context, course *models.Course) error {

	err := u.courseRepo.CreateCourse(ctx, course)

	return err

}

func (u *courseUsecase) GetCourseByCode(ctx context.Context, code string) (*models.Course, error) {

	course, err := u.courseRepo.GetCourseByCode(ctx, code)

	if err != nil {
		return nil, errors.New("course not found")
	}

	for _, resourceID := range course.ResourceIDs {

		log.Println(resourceID)

		resource, err := resource.ResourceUsecase.GetResourceByID(ctx, resourceID)

		if err != nil {
			return nil, err
		}

		course.Resources = append(course.Resources, resource)

	}

	return course, nil
}
