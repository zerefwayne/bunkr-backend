package course

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

type Repository interface {
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCourseByCode(ctx context.Context, code string) (*models.Course, error)
	CreateCourse(ctx context.Context, course *models.Course) error
	UpdateCourse(ctx context.Context, course *models.Course) error
}
