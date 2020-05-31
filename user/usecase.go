package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zerefwayne/college-portal-backend/models"
	"github.com/zerefwayne/college-portal-backend/utils"
)

var ctx context.Context = context.Background()

// Usecase Defines the usecase for User
type Usecase interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type userUsecase struct {
	userRepo Repository
}

var UserUsecase userUsecase

func (u *userUsecase) GetByID(ctx context.Context, id string) (*models.User, error) {

	user, err := u.userRepo.GetByID(ctx, id)

	return user, err
}

func (u *userUsecase) GetByUsername(ctx context.Context, username string) (*models.User, error) {

	user, err := u.userRepo.GetByUsername(ctx, username)

	return user, err
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	user, err := u.userRepo.GetByEmail(ctx, email)

	return user, err
}

func (u *userUsecase) CreateUser(ctx context.Context, user *models.User) error {

	// Set new user ID
	user.ID = uuid.New().String()

	if _, err := u.GetByEmail(ctx, user.Email); err == nil {
		return errors.New("email already exists")
	}

	if _, err := u.GetByUsername(ctx, user.Username); err == nil {
		return errors.New("username already exists")
	}

	hash, err := utils.GenerateHashFromPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hash

	u.userRepo.CreateUser(ctx, user)

	return nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
