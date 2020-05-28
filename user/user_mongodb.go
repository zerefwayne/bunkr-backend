package user

import (
	"context"
	"fmt"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	DB *mongo.Database
}

// NewMongoUserRepository ...
func NewMongoUserRepository(client *mongo.Client) Repository {
	database := client.Database(config.C.Env.Database)

	return &userRepository{
		DB: database,
	}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return nil, nil
}
func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {

	fmt.Printf("mongodb %+v\n", user)

	result, err := r.DB.Collection("user").InsertOne(ctx, user)

	if err != nil {
		return err
	}

	fmt.Printf("successfully written! %+v\n", result)

	return nil
}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
