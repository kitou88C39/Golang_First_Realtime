package customwebsocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResposeWriter, r *http.Request){
	upgrader.CheckOrigin = func (r *http.Request)  {
		
	}
}