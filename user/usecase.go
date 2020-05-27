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

type userUsecase struct {
	userRepo Repository
}

// NewUserUsecase ...
func NewUserUsecase(r Repository) Usecase {
	return &userUsecase{
		userRepo: r,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}
func (u *userUsecase) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}
func (u *userUsecase) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
