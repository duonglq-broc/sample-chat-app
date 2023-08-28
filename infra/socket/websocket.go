package socket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WS struct {
	Conn *websocket.Conn
}

func InitWS(c echo.Context) (*WS, error) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return nil, err
	}

	return &WS{
		Conn: conn,
	}, err
}

func (w *WS) CLose() error {
	err := w.Conn.Close()
	if err != nil {
		return err
	}

	return nil
}
