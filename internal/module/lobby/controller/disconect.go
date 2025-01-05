package controller

import (
	"context"
	"fmt"

	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (lc *LobbyController) Disconect(ctx context.Context, c *websocket.Conn, msg []byte) {
	gameId := lc.WebSocketStorage.ConnGameId[c]
	userId := lc.WebSocketStorage.ConnUserId[c]

	modelUser, err := lc.userService.GetUserById(ctx, userId)
	if err != nil {
		log.Errorf("failed to get user by id, error: %v", err)
		c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusInternalServerError,
			},
		})
		return
	}

	for _, conn := range lc.WebSocketStorage.GameIdConn[gameId] {
		conn.WriteJSON(response.Respons{
			Header: response.Header{
				Code:  fiber.StatusOK,
				Event: "disconect",
			},
			Body: bson.M{
				"message": fmt.Sprintf("%s, потеря соединения.", modelUser.Username),
				"data": bson.M{
					"id": modelUser.Id,
				},
			},
		})
	}

	if err := lc.WebSocketStorage.DeleteConn(c); err != nil {
		log.Errorf("failed to delete conn, error: %v", err)
		return
	}
}
