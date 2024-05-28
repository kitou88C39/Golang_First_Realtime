package main

import (
	customwebsocket "chatapplication/websocket"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool,w http.ResponseWriter, r *http.Request) {
    customwebsocket.Upgrade(w,r)
}

func setupRoutes(){
    pool := customwebsocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serverWs(w, r, pool)
    })
}

func main() {
    setupRoutes()
    http.ListenAndServe(":9000", nil)
};