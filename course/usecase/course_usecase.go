package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/course"
	"github.com/zerefwayne/college-portal-backend/course/repository"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/resource"
	"github.com/zerefwayne/college-portal-backend/resource/usecase"
)

type courseUsecase struct {
	courseRepo course.Repository
	resource   resource.Usecase
}

func NewCourseUsecase() course.Usecase {

	return &courseUsecase{
		courseRepo: repository.NewMongoResourceRepository(config.C.MongoDB),
		resource:   usecase.NewResourceUsecase(),
	}

}

var CourseUsecase courseUsecase

func (u *courseUsecase) PushResource(ctx context.Context, courseCode string, resourceID string) error {

	course, err := u.GetCourseByCode(ctx, courseCode)

	if err != nil {
		return err
	}

	course.ResourceIDs = append(course.ResourceIDs, resourceID)

	if err := u.UpdateCourse(ctx, course); err != nil {
		return err
	}

	return nil

}

func (u *courseUsecase) UpdateCourse(ctx context.Context, course *models.Course) error {

	err := u.courseRepo.UpdateCourse(ctx, course)

	return err

}

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

		resource, err := u.resource.GetResourceByID(ctx, resourceID)

		if err != nil {
			return nil, err
		}

		course.Resources = append(course.Resources, resource)

	}

	return course, nil
}
