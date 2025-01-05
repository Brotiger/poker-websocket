package service

import (
	"context"
	"fmt"

	cError "github.com/Brotiger/poker-websocket/internal/module/lobby/error"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectTokenService struct {
	connectTokenRepository *repository.ConnectTokenRepository
}

func NewConnectTokenService() *ConnectTokenService {
	return &ConnectTokenService{
		connectTokenRepository: repository.NewConnectTokenRepository(),
	}
}

type RequestVerifyTokenDTO struct {
	GameId primitive.ObjectID
	UserId primitive.ObjectID
	Token  string
}

func (cts *ConnectTokenService) VerifyToken(ctx context.Context, requestVerifyTokenDTO RequestVerifyTokenDTO) (bool, error) {
	modelToken, err := cts.connectTokenRepository.FindTokenByToken(ctx, requestVerifyTokenDTO.Token)
	if err != nil {
		return false, fmt.Errorf("failed to find token by token, error: %w", err)
	}

	if modelToken.GameId != requestVerifyTokenDTO.GameId {
		return false, cError.ErrInvalidConnectToken
	}

	if modelToken.UserId != requestVerifyTokenDTO.UserId {
		return false, cError.ErrInvalidConnectToken
	}

	if err := cts.connectTokenRepository.DeleteTokenById(ctx, modelToken.Id); err != nil {
		return false, fmt.Errorf("failed to delete token by id, error: %w", err)
	}

	return true, nil
}
