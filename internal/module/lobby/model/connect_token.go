package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConnectToken struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"userId"`
	GameId    primitive.ObjectID `bson:"gameId"`
	Token     string             `bson:"token"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	CreatedAt time.Time          `bson:"createdAt"`
}
