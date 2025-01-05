package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	coreApiService "github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/Brotiger/poker-websocket/internal/config"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/request"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/service"
	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/Brotiger/poker-websocket/internal/storage"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (lc *LobbyController) Join(ctx context.Context, c *websocket.Conn, msg []byte) {
	var requestJoin request.Join
	if err := json.Unmarshal(msg, &requestJoin); err != nil {
		log.Errorf("failed to unmarshal json, error: %v", err)
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusInternalServerError,
			},
		})
		return
	}

	token, err := lc.tokenService.GetToken(requestJoin.Header.AccessToken)
	if err != nil {
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusBadRequest,
			},
			Body: bson.M{
				"message": "Неверный формат токена.",
			},
		})
	}

	tokenClaims, err := lc.tokenService.VerifyToken(token, config.Cfg.JWT.Secret)
	if err != nil {
		if errors.Is(err, coreApiService.ErrInvalidToken) {
			c.WriteJSON(response.Respons{
				Header: response.Header{
					Code: fiber.StatusUnauthorized,
				},
				Body: bson.M{
					"message": "Просроченный токен доступа.",
				},
			})
			return
		}

		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusBadRequest,
			},
			Body: bson.M{
				"message": "Невалидный токен.",
			},
		})
		return
	}

	lc.ConnectTokenService.VerifyToken(ctx, service.RequestVerifyTokenDTO{
		Token:  requestJoin.Header.ConnectToken,
		GameId: requestJoin.Body.GameId,
		UserId: tokenClaims.UserId,
	})

	for client, userId := range storage.WebSocketClients[requestJoin.Body.GameId] {
		modelUser, err := lc.userService.GetUserById(ctx, userId)
		if err != nil {
			log.Errorf("failed to get user by id, error: %v", err)
			continue
		}

		client.WriteJSON(response.Respons{
			Header: response.Header{
				Code:  fiber.StatusOK,
				Event: "join",
			},
			Body: bson.M{
				"message": fmt.Sprintf("Игрок %s подключился к игре", modelUser.Username),
			},
		})
	}

	storage.WebSocketClients[requestJoin.Body.GameId][c] = tokenClaims.UserId
}
