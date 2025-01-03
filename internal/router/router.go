package router

import (
	"encoding/json"
	"fmt"

	"github.com/Brotiger/poker-websocket/internal/model"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/controller"
	"github.com/gofiber/contrib/websocket"
)

type Router struct {
	lobbyController *controller.LobbyController
}

func NewRouter() *Router {
	return &Router{
		lobbyController: controller.NewLobbyController(),
	}
}

func (r *Router) ProcessMessage(c *websocket.Conn) error {
	_, msg, err := c.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read message, error: %w", err)
	}

	var modelMessage model.Message
	if err := json.Unmarshal(msg, &modelMessage); err != nil {
		return fmt.Errorf("failed to unmarshal message, error: %w", err)
	}

	switch modelMessage.Type {
	case "join":
		if err := r.lobbyController.Join(c); err != nil {
			return fmt.Errorf("failed to join, error: %w", err)
		}
	}

	return nil
}
