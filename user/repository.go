package user

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/models"
)

// Repository defines the usecase for User
type Repository interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type userRepository struct {
}

// NewUserRepository ...
func NewUserRepository() Repository {
	return &userUsecase{}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
