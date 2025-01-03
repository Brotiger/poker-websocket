package controller

import "github.com/Brotiger/poker-websocket/internal/middleware"

type LobbyController struct {
	authMiddleware *middleware.AuthMiddleware
}

func NewLobbyController() *LobbyController {
	return &LobbyController{
		authMiddleware: middleware.NewAuthMiddleware(),
	}
}
