package controller

import (
	pkgService "github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/service"
)

type LobbyController struct {
	tokenService        *pkgService.TokenService
	userService         *service.UserService
	ConnectTokenService *service.ConnectTokenService
}

func NewLobbyController() *LobbyController {
	return &LobbyController{
		tokenService:        pkgService.NewTokenService(),
		userService:         service.NewUserService(),
		ConnectTokenService: service.NewConnectTokenService(),
	}
}
