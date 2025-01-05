package router

import (
	"context"
	"encoding/json"

	"github.com/Brotiger/poker-websocket/internal/module/lobby/controller"
	"github.com/Brotiger/poker-websocket/internal/request"
	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type Router struct {
	lobbyController *controller.LobbyController
}

func NewRouter() *Router {
	return &Router{
		lobbyController: controller.NewLobbyController(),
	}
}

func (r *Router) ProcessMessage(ctx context.Context, c *websocket.Conn) {
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Errorf("failed to read message, error: %v", err)
		if err := c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusInternalServerError,
			},
		}); err != nil {
			log.Errorf("failed to write response, error: %v", err)
		}
		return
	}

	var requestMessage request.Message
	if err := json.Unmarshal(msg, &requestMessage); err != nil {
		log.Errorf("failed to unmarshal message, error: %v", err)
		if err := c.WriteJSON(response.Respons{
			Header: response.Header{
				Code: fiber.StatusInternalServerError,
			},
		}); err != nil {
			log.Errorf("failed to write response, error: %v", err)
		}
		return
	}

	switch requestMessage.Event {
	case "join":
		r.lobbyController.Join(ctx, c, msg)
		return
	}

	if err := c.WriteJSON(response.Respons{
		Header: response.Header{
			Code: fiber.StatusNotFound,
		},
		Body: bson.M{
			"message": "Неизвестный тип ивента.",
		},
	}); err != nil {
		log.Errorf("failed to write response, error: %v", err)
	}
}
