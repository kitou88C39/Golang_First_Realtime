package main

import (
	customwebsocket "chatapplication/websocket"
	"log"
	"net/http"
)

func serverWs(pool *customwebsocket.Pool,w http.ResponseWriter, r *http.Request) {
    conn, err := customwebsocket.Upgrade(w,r)
    if err != nil {
        log.Println(err)
        return
    }
}

func setupRoutes(){
    pool := customwebsocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serverWs(pool, w, r)
    })
}

func main() {
    setupRoutes()
    http.ListenAndServe(":9000", nil)
};