package main
import "net/http"

func serverWs(w http.ResposeWriter, r *http.Request){}

func main() {
    http.ListenAndServe(":9000", nil)
};