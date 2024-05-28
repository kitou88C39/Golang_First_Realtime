package main

import (
	"net/http"
)

func serverWs(w http.ResponseWriter, r *http.Request) {
    // WebSocketの接続処理をここに書く
}

func setupRoutes(){
    
}

func main() {
    setupRoutes()
    http.ListenAndServe(":9000", nil)
};