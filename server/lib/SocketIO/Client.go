package SocketIO

import "github.com/gorilla/websocket"

type Client struct {
	Conn  *websocket.Conn
	Emits []SocketMessageEvent
}

func NewClient(conn *websocket.Conn) Client {
	return Client{
		Conn:  conn,
		Emits: make([]SocketMessageEvent, 0),
	}
}
