package course

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
)

type Usecase interface {
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCourseByCode(ctx context.Context, code string) (*models.Course, error)
}

type courseUsecase struct {
}

// NewCourseUsecase ...
func NewCourseUsecase() Usecase {
	return &courseUsecase{}
}

func (u *courseUsecase) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	courses := config.Courses
	return courses, nil
}

func (u *courseUsecase) GetCourseByCode(ctx context.Context, code string) (*models.Course, error) {
	return nil, nil
}
