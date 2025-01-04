package main

import (
	"context"
	"time"

	mongodb "github.com/Brotiger/poker-core_api/pkg/mongodb/connection"
	"github.com/Brotiger/poker-websocket/internal/config"
	"github.com/Brotiger/poker-websocket/internal/connection"
	"github.com/Brotiger/poker-websocket/internal/router"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.Cfg.Fiber.DisableStartupMessage,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	mongodbCtx, cancelMongodbCtx := context.WithTimeout(ctx, time.Duration(config.Cfg.MongoDB.ConnectTimeoutMs)*time.Millisecond)
	defer cancelMongodbCtx()

	mongodbClient, err := mongodb.Connect(
		mongodbCtx,
		config.Cfg.MongoDB.Uri,
		config.Cfg.MongoDB.Username,
		config.Cfg.MongoDB.Password,
	)
	if err != nil {
		log.Fatalf("failed to connect to mongodb, error: %v", err)
	}

	connection.DB = mongodbClient.Database(config.Cfg.MongoDB.Database)

	router := router.NewRouter()
	app.Post("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			log.Info("Connection closed")
			c.Close()
		}()

		log.Info("WebSocket connection opened")

		for {
			router.ProcessMessage(ctx, c)
		}
	}))

	log.Fatal(app.Listen(config.Cfg.Fiber.Listen))
}
