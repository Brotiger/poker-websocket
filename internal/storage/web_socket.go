package storage

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @todo добавить мьютекс
type WebSocketStorage struct {
	GameIdConnMap map[*websocket.Conn]int
	GameIdConn    map[primitive.ObjectID][]*websocket.Conn
	ConnGameId    map[*websocket.Conn]primitive.ObjectID
	ConnUserId    map[*websocket.Conn]primitive.ObjectID
}

func NewWebSockeStorage() *WebSocketStorage {
	return &WebSocketStorage{
		GameIdConn: make(map[primitive.ObjectID][]*websocket.Conn),
		ConnGameId: make(map[*websocket.Conn]primitive.ObjectID),
		ConnUserId: make(map[*websocket.Conn]primitive.ObjectID),
	}
}

type RequestAddConn struct {
	GameId primitive.ObjectID
	UserId primitive.ObjectID
	Socket *websocket.Conn
}

func (wss *WebSocketStorage) AddConn(requestAddConn RequestAddConn) {
	wss.GameIdConn[requestAddConn.GameId] = append(wss.GameIdConn[requestAddConn.GameId], requestAddConn.Socket)
	wss.GameIdConnMap[requestAddConn.Socket] = len(wss.GameIdConn[requestAddConn.GameId])
	wss.ConnGameId[requestAddConn.Socket] = requestAddConn.GameId
	wss.ConnUserId[requestAddConn.Socket] = requestAddConn.UserId
}

func (wss *WebSocketStorage) DeleteConn(c *websocket.Conn) error {
	delete(wss.ConnGameId, c)
	delete(wss.ConnUserId, c)
	index, ok := wss.GameIdConnMap[c]
	if !ok {
		return fmt.Errorf("failed to get conn map by game id")
	}

	gameId, ok := wss.ConnGameId[c]
	if !ok {
		return fmt.Errorf("failed to get game id by conn")
	}

	conn, ok := wss.GameIdConn[gameId]
	if !ok {
		return fmt.Errorf("failed to get conn by game id")
	}

	conn[len(conn)-1], conn[index] = conn[index], conn[len(conn)-1]

	return nil
}
