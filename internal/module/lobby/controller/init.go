package controller

import (
	pkgService "github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/service"
	"github.com/Brotiger/poker-websocket/internal/storage"
)

type LobbyController struct {
	tokenService        *pkgService.TokenService
	userService         *service.UserService
	ConnectTokenService *service.ConnectTokenService
	WebSocketStorage    *storage.WebSocketStorage
}

func NewLobbyController() *LobbyController {
	return &LobbyController{
		tokenService:        pkgService.NewTokenService(),
		userService:         service.NewUserService(),
		ConnectTokenService: service.NewConnectTokenService(),
		WebSocketStorage:    storage.NewWebSockeStorage(),
	}
}
