package router

import (
	"encoding/json"

	"github.com/Brotiger/poker-websocket/internal/model"
	"github.com/Brotiger/poker-websocket/internal/module/lobby/controller"
	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/gofiber/contrib/websocket"
	log "github.com/sirupsen/logrus"
)

type Router struct {
	lobbyController *controller.LobbyController
}

func NewRouter() *Router {
	return &Router{
		lobbyController: controller.NewLobbyController(),
	}
}

func (r *Router) ProcessMessage(c *websocket.Conn) {
	var res response.Response
	res = &response.InternalServerError{}

	defer func() {
		res.Send(c)
	}()

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Errorf("failed to read message, error: %v", err)
		return
	}

	var modelMessage model.Message
	if err := json.Unmarshal(msg, &modelMessage); err != nil {
		log.Errorf("failed to unmarshal message, error: %v", err)
		return
	}

	switch modelMessage.Type {
	case "join":
		res = r.lobbyController.Join(c, msg)
	}

	log.Errorf("undefined message type, error: %v", err)
}
