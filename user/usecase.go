package user

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/zerefwayne/college-portal-backend/common"
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
	AddCourse(ctx context.Context, userID string, courseCode string) error
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetSubscribedCourses(ctx context.Context, id string) ([]*models.Course, error)
}

type userUsecase struct {
	userRepo Repository
}

var UserUsecase userUsecase

func (u *userUsecase) AddCourse(ctx context.Context, userID string, courseCode string) error {

	user, err := UserUsecase.GetByID(ctx, userID)

	if err != nil {
		return err
	}

	log.Println("User loaded", user)

	course, err := common.Course.GetCourseByCode(ctx, courseCode)

	if err != nil {
		return err
	}

	log.Println("Course loaded", course)

	err = u.userRepo.AddCourse(ctx, userID, courseCode)

	return err

}

func (u *userUsecase) GetSubscribedCourses(ctx context.Context, id string) ([]*models.Course, error) {

	user, err := u.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	log.Println("Fetching courses for", user.Username)

	var courses []*models.Course

	for _, courseCode := range user.SubscribedCourses {

		course, err := common.Course.GetCourseByCode(ctx, courseCode)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)

	}

	return courses, nil

}

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
	user.SubscribedCourses = []string{}

	u.userRepo.CreateUser(ctx, user)

	return nil
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
