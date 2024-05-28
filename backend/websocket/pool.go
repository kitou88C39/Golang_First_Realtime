package customwebsocket

type pool struct{
	Register chan *Client
	Unregister chan *Client
	Clients map [*Client] pool
}

type Message struct {
	Type int `json:"type"`
	Body string `json:body`

}