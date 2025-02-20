package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Join struct {
	Header struct {
		AccessToken  string `json:"access_token"`
		ConnectToken string `json:"connect_token"`
	} `json:"header"`
	Body struct {
		GameId primitive.ObjectID `json:"game_id"`
	} `json:"body"`
}
