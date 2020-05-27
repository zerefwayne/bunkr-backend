package user

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Usecase Defines the usecase for User
type Usecase interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}
