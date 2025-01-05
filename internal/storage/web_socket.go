package storage

import (
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @todo добавить мьютекс
type WebSocketStorage struct {
	connectionInGameMap map[*websocket.Conn]int
	gameIdConnections   map[primitive.ObjectID][]*websocket.Conn

	connGameId map[*websocket.Conn]primitive.ObjectID
	connUserId map[*websocket.Conn]primitive.ObjectID
	mux        sync.Mutex
}

func NewWebSockeStorage() *WebSocketStorage {
	return &WebSocketStorage{
		connectionInGameMap: make(map[*websocket.Conn]int),
		gameIdConnections:   make(map[primitive.ObjectID][]*websocket.Conn),

		connGameId: make(map[*websocket.Conn]primitive.ObjectID),
		connUserId: make(map[*websocket.Conn]primitive.ObjectID),
	}
}

type RequestAddConn struct {
	GameId     primitive.ObjectID
	UserId     primitive.ObjectID
	Connection *websocket.Conn
}

func (wss *WebSocketStorage) AddConn(requestAddConn RequestAddConn) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	wss.gameIdConnections[requestAddConn.GameId] = append(wss.gameIdConnections[requestAddConn.GameId], requestAddConn.Connection)
	wss.connectionInGameMap[requestAddConn.Connection] = len(wss.gameIdConnections[requestAddConn.GameId]) - 1

	wss.connGameId[requestAddConn.Connection] = requestAddConn.GameId
	wss.connUserId[requestAddConn.Connection] = requestAddConn.UserId
}

func (wss *WebSocketStorage) DeleteConn(c *websocket.Conn) error {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	connectionInGameIndex, ok := wss.connectionInGameMap[c]
	if !ok {
		return fmt.Errorf("failed to get conn map by game id")
	}

	gameId, ok := wss.connGameId[c]
	if !ok {
		return fmt.Errorf("failed to get game id by conn")
	}

	gameConnections, ok := wss.gameIdConnections[gameId]
	if !ok {
		return fmt.Errorf("failed to get conn by game id")
	}

	delete(wss.connGameId, c)
	delete(wss.connUserId, c)

	gameConnections[len(gameConnections)-1], gameConnections[connectionInGameIndex] = gameConnections[connectionInGameIndex], gameConnections[len(gameConnections)-1]

	return nil
}

func (wss *WebSocketStorage) GetConnByGameId(gameId primitive.ObjectID) ([]*websocket.Conn, error) {
	wss.mux.Lock()
	defer wss.mux.Unlock()

	connection, ok := wss.gameIdConnections[gameId]
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
