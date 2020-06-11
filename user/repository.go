package user

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Repository defines the usecase for User
type Repository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	AddCourse(ctx context.Context, userID string, courseCode string) error
	RemoveCourse(ctx context.Context, userID string, courseCode string) error
	AddBookmark(ctx context.Context, userID string, resourceID string) error
	RemoveBookmark(ctx context.Context, userID string, resourceID string) error
}
