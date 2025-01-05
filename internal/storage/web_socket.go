package storage

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @todo добавить мьютекс
type WebSocketStorage struct {
	gameIdConnMap map[*websocket.Conn]int
	gameIdConn    map[primitive.ObjectID][]*websocket.Conn
	connGameId    map[*websocket.Conn]primitive.ObjectID
	connUserId    map[*websocket.Conn]primitive.ObjectID
	mux           sync.Mutex
}

func NewWebSockeStorage() *WebSocketStorage {
	return &WebSocketStorage{
		gameIdConn: make(map[primitive.ObjectID][]*websocket.Conn),
		connGameId: make(map[*websocket.Conn]primitive.ObjectID),
		connUserId: make(map[*websocket.Conn]primitive.ObjectID),
	}
}

type RequestAddConn struct {
	GameId primitive.ObjectID
	UserId primitive.ObjectID
	Socket *websocket.Conn
}

func (wss *WebSocketStorage) AddConn(requestAddConn RequestAddConn) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	wss.gameIdConn[requestAddConn.GameId] = append(wss.gameIdConn[requestAddConn.GameId], requestAddConn.Socket)
	wss.gameIdConnMap[requestAddConn.Socket] = len(wss.gameIdConn[requestAddConn.GameId])
	wss.connGameId[requestAddConn.Socket] = requestAddConn.GameId
	wss.connUserId[requestAddConn.Socket] = requestAddConn.UserId
}

func (wss *WebSocketStorage) DeleteConn(c *websocket.Conn) error {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	index, ok := wss.gameIdConnMap[c]
	if !ok {
		return fmt.Errorf("failed to get conn map by game id")
	}

	gameId, ok := wss.connGameId[c]
	if !ok {
		return fmt.Errorf("failed to get game id by conn")
	}

	conn, ok := wss.gameIdConn[gameId]
	if !ok {
		return fmt.Errorf("failed to get conn by game id")
	}

	delete(wss.connGameId, c)
	delete(wss.connUserId, c)

	conn[len(conn)-1], conn[index] = conn[index], conn[len(conn)-1]

	return nil
}

func (wss *WebSocketStorage) GetConnByGameId(gameId primitive.ObjectID) ([]*websocket.Conn, error) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	connection, ok := wss.gameIdConn[gameId]
	if !ok {
		return nil, fmt.Errorf("failed to get conn by game id")
	}

	return connection, nil
}

func (wss *WebSocketStorage) GetGameIdByConn(connection *websocket.Conn) (*primitive.ObjectID, error) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	gameId, ok := wss.connGameId[connection]
	if !ok {
		return nil, fmt.Errorf("failed to get game id by conn")
	}

	return &gameId, nil
}

func (wss *WebSocketStorage) GetUserIdByConn(connection *websocket.Conn) (*primitive.ObjectID, error) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	userId, ok := wss.connUserId[connection]
	if !ok {
		return nil, fmt.Errorf("failed to get user id by conn")
	}

	return &userId, nil
}
