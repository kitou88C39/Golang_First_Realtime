package customwebsocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type pool struct{
	Conn * websocket.Conn
	Pool *Pool
	mu sync.Mutex
}

type Message struct {
	Type int `json:"type"`
	Body string `json:body`

}