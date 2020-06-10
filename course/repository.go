package course

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

type Repository interface {
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCourseByCode(ctx context.Context, code string) (*models.Course, error)
	GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error)
	CreateCourse(ctx context.Context, course *models.Course) error
	PushResource(ctx context.Context, courseCode string, resourceID string) error
	PopResource(ctx context.Context, resourceID string) error
}
