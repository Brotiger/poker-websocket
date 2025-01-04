package storage

import (
	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var WebSocketClients = make(map[primitive.ObjectID]map[*websocket.Conn]primitive.ObjectID)
