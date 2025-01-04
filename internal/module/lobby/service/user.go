package service

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-websocket/internal/module/lobby/model"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) GetUserById(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	modelUser, err := us.userRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id, error %w", err)
	}

	return modelUser, nil
}
