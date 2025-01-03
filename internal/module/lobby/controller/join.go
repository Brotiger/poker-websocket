package controller

import (
	"encoding/json"

	"github.com/Brotiger/poker-websocket/internal/module/lobby/model"
	"github.com/Brotiger/poker-websocket/internal/response"
	"github.com/gofiber/contrib/websocket"
	log "github.com/sirupsen/logrus"
)

func (lc *LobbyController) Join(c *websocket.Conn, msg []byte) response.Response {
	var modelJoin model.Join
	if err := json.Unmarshal(msg, &modelJoin); err != nil {
		log.Errorf("failed to unmarshal json, error: %v", err)
		return &response.InternalServerError{}
	}

	res, _ := lc.authMiddleware.VerifyToken(modelJoin.Header.AccessToken)
	if res != nil {
		return res
	}

	return &response.OK{}
}
