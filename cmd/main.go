package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"sample-chat-app/handler"
	"sample-chat-app/infra/db"
)

func main() {
	e := echo.New()

	ctx := context.Background()

	dbClient, err := db.InitDB(ctx, "mongodb://localhost:27017")
	defer dbClient.Close(ctx)

	e.GET("/", func(c echo.Context) error {
		return c.File("public/index.html")
	})

	e.GET("/ws", func(c echo.Context) error {
		err := handler.WebSocketHandler(c)
		if err != nil {
			return err
		}
		return nil
	})

	err = e.Start(":8080")
	if err != nil {
		return
	}
}
