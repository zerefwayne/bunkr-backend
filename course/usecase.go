package course

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource"
)

type Usecase interface {
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCourseByCode(ctx context.Context, code string) (*models.Course, error)
}

type courseUsecase struct {
}

var CourseUsecase courseUsecase

func (u *courseUsecase) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	courses := config.Courses
	return courses, nil
}

func (u *courseUsecase) GetCourseByCode(ctx context.Context, code string) (*models.Course, error) {

	var course *models.Course

	for _, c := range config.Courses {

		if c.Code == code {
			course = c
			break
		}

	}

	if course == nil {
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
