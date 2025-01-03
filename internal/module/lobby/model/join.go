package model

import (
	"github.com/Brotiger/poker-websocket/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Join struct {
	model.Message
	Body struct {
		GameId primitive.ObjectID `json:"game_id"`
	} `json:"body"`
}
