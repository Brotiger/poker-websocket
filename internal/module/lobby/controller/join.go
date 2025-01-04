package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/request"
	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/Brotiger/poker-websocket/internal/storage"
	"github.com/gofiber/contrib/websocket"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lc *LobbyController) Join(ctx context.Context, c *websocket.Conn, msg []byte) {
	var requestJoin request.Join
	if err := json.Unmarshal(msg, &requestJoin); err != nil {
		log.Errorf("failed to unmarshal json, error: %v", err)
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: 500,
			},
		})
		return
	}

	token, err := lc.tokenService.GetToken(requestJoin.Header.AccessToken)
	if err != nil {
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: 400,
			},
			Body: bson.M{
				"message": "Неверный формат токена.",
			},
		})
	}

	tokenClaims, err := lc.tokenService.VerifyToken(token)
	if err != nil {
		if errors.Is(err, service.ErrInvalidToken) {
			c.WriteJSON(response.Respons{
				Header: response.Header{
					Code: 401,
				},
				Body: bson.M{
					"message": "Просроченный токен доступа.",
				},
			})
			return
		}

		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: 400,
			},
			Body: bson.M{
				"message": "Невалидный токен.",
			},
		})
		return
	}

	userId, err := primitive.ObjectIDFromHex(tokenClaims.UserId)
	if err != nil {
		log.Errorf("failed to convert user id to object id, error: %v", err)
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: 500,
			},
		})
		return
	}

	storage.WebSocketClients[requestJoin.Body.GameId][c] = userId

	for client, userId := range storage.WebSocketClients[requestJoin.Body.GameId] {
		modelUser, err := lc.userService.GetUserById(ctx, userId)
		if err != nil {
			log.Errorf("failed to get user by id, error: %v", err)
			continue
		}

		client.WriteJSON(response.Respons{
			Header: response.Header{
				Code:  200,
				Event: "join",
			},
			Body: bson.M{
				"message": fmt.Sprintf("Игрок %s подключился к игре", modelUser.Username),
			},
		})
	}
}
