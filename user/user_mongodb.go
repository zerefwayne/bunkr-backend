package user

import (
	"context"

	"github.com/zerefwayne/college-portal-backend/config"
	"github.com/zerefwayne/college-portal-backend/models"
	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"username": username}

	var user models.User

	err := r.DB.Collection("user").FindOne(ctx, filter).Decode(&user)

	return &user, err
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {

	filter := bson.M{"email": email}

	var user models.User

	err := r.DB.Collection("user").FindOne(ctx, filter).Decode(&user)

	return &user, err
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {

	_, err := r.DB.Collection("user").InsertOne(ctx, user)

	return err

}
func (r *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return nil, nil
}
