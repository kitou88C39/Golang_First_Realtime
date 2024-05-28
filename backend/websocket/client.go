package customwebsocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct{
	Conn * websocket.Conn
	Pool *Pool
	mu sync.Mutex
}

type Message struct {
    Type int `json:"type"`
    Body string `json:"body"`
}

func(c *Pool) Start(){
	defer func ()  {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
					return

		}
	}
}