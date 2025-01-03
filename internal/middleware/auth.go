package middleware

import (
	"errors"
	"strings"

	"github.com/Brotiger/poker-core_api/pkg/mongodb/model"
	"github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/Brotiger/poker-websocket/internal/response"
)

const headerPrefix = "Bearer"

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}

func (am *AuthMiddleware) VerifyToken(accessToken string) (response.Response, *model.JWTClaims) {
	token, err := am.getToken(accessToken)
	if err != nil {
		return &response.BadRequest{
			Message: "Неверный формат токена.",
		}, nil
	}

	tokenClaims, err := am.tokenService.VerifyToken(token)
	if err != nil {
		return &response.BadRequest{
			Message: "Невалидный токен.",
		}, nil
	}

	return nil, tokenClaims
}

func (am *AuthMiddleware) getToken(accessToken string) (string, error) {
	l := len(headerPrefix)
	if len(accessToken) < l+2 || accessToken[:l] != headerPrefix {
		return "", errors.New("invalid token format")
	}

	return strings.TrimSpace(accessToken[l:]), nil
}
