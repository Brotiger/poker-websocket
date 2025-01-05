package repository

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-websocket/internal/config"
	"github.com/Brotiger/poker-websocket/internal/connection"
	cError "github.com/Brotiger/poker-websocket/internal/module/lobby/error"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConnectTokenRepository struct{}

func NewConnectTokenRepository() *ConnectTokenRepository {
	return &ConnectTokenRepository{}
}

func (ctr *ConnectTokenRepository) FindTokenByToken(ctx context.Context, token string) (*model.ConnectToken, error) {
	var modelConnectToken model.ConnectToken
	if err := connection.DB.Collection(config.Cfg.MongoDB.Table.ConnectToken).FindOne(
		ctx,
		bson.M{
			"token": token,
		},
	).Decode(&modelConnectToken); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, cError.ErrConnectTokenNotFound
		}

		return nil, fmt.Errorf("failed to find one, error: %w", err)
	}

	return &modelConnectToken, nil
}

func (ctr *ConnectTokenRepository) DeleteTokenById(ctx context.Context, id primitive.ObjectID) error {
	if _, err := connection.DB.Collection(config.Cfg.MongoDB.Table.ConnectToken).DeleteOne(
		ctx,
		bson.M{
			"_id": id,
		},
	); err != nil {
		return fmt.Errorf("failed to delete one, error: %w", err)
	}

	return nil
}
