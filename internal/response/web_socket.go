package response

import (
	"encoding/json"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2/log"
)

type Response interface {
	Send(c *websocket.Conn)
}

type BadRequest struct {
	Message string
	Errors  any
}

func (br *BadRequest) Send(c *websocket.Conn) {
	type Body struct {
		Message string `json:"message"`
		Errors  any    `json:"errors,omitempty"`
	}

	type BadRequest struct {
		Code int  `json:"code"`
		Body Body `json:"body"`
	}

	responseBadRequest, err := json.Marshal(BadRequest{
		Code: 400,
		Body: Body{
			Message: br.Message,
			Errors:  br.Errors,
		},
	})
	if err != nil {
		log.Errorf("failed to marshal, error: %w", err)
		c.WriteMessage(8, nil)

		return
	}

	c.WriteMessage(1, responseBadRequest)
}

type NotFound struct {
	Message string
}

func (nf *NotFound) Send(c *websocket.Conn) {
	type Body struct {
		Message string `json:"message"`
	}

	type NotFound struct {
		Code int  `json:"code"`
		Body Body `json:"body"`
	}

	responseNotFound, err := json.Marshal(NotFound{
		Code: 404,
		Body: Body{
			Message: nf.Message,
		},
	})
	if err != nil {
		log.Errorf("failed to marshal, error: %w", err)
		c.WriteMessage(8, nil)

		return
	}

	c.WriteMessage(1, responseNotFound)
}

type InternalServerError struct{}

func (ise *InternalServerError) Send(c *websocket.Conn) {
	type InternalServerError struct {
		Code int `json:"code"`
	}

	responseInternalServerError, err := json.Marshal(InternalServerError{
		Code: 500,
	})
	if err != nil {
		log.Errorf("failed to marshal, error: %w", err)
		c.WriteMessage(8, nil)

		return
	}

	c.WriteMessage(1, responseInternalServerError)
}

type Unauthorized struct{}

func (u *Unauthorized) Send(c *websocket.Conn) {
	type Unauthorized struct {
		Code int `json:"code"`
	}

	responseUnauthorized, err := json.Marshal(Unauthorized{
		Code: 401,
	})
	if err != nil {
		log.Errorf("failed to marshal, error: %w", err)
		c.WriteMessage(8, nil)

		return
	}

	c.WriteMessage(1, responseUnauthorized)
}

type OK struct {
	Body any
}

func (o *OK) Send(c *websocket.Conn) {
	type OK struct {
		Code int `json:"code"`
		Body any `json:"body,omitempty"`
	}

	responseOK, err := json.Marshal(OK{
		Code: 200,
		Body: o.Body,
	})
	if err != nil {
		log.Errorf("failed to marshal, error: %w", err)
		c.WriteMessage(8, nil)

		return
	}

	c.WriteMessage(1, responseOK)
}
