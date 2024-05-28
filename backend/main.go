package main

import (
	customwebsocket "chatapplication/websocket"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool,w http.ResponseWriter, r *http.Request) {
    // WebSocketの接続処理をここに書く
}

func setupRoutes(){
    pool := customwebsocket.NewPool()
    go pool.Start()
    http.HandlerFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serverWs(pool, w, r)
    })
}

func main() {
    setupRoutes()
    http.ListenAndServe(":9000", nil)
};