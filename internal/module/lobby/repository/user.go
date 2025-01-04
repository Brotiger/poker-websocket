package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-websocket/internal/config"
	"github.com/Brotiger/poker-websocket/internal/connection"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindUserById(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	var modelUser model.User
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.User).FindOne(
		ctx,
		bson.M{
			"_id": id,
		},
	).Decode(&modelUser); err != nil {
		if err == mongo.ErrNoDocuments {

		}

		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelUser, nil
}
