package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"sample-chat-app/infra/socket"
)

func WebSocketHandler(c echo.Context) error {
	ws, err := socket.InitWS(c)
	if err != nil {
		return err
	}

	defer ws.CLose()

	for {
		messageType, p, err := ws.Conn.ReadMessage()
		if err != nil {
			return err
		}

		fmt.Println(messageType, string(p))
	}
}
