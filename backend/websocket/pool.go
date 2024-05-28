package customwebsocket

import "fmt"

type Pool struct{
	Register chan *Client
	Unregister chan *Client
	Clients map [*Client] bool
	Broadcast chan Message
}

func NewPool() *Pool{return &Pool{
	Register: make(chan *Client),
	Unregister:  make(chan *Client),
	Clients: make(map [*Client] bool),
	Broadcast: make(chan Message),
}}

func(poll *Pool) Start(){
	for {
		select{
		case client := <-poll.Register:
			pool.Clients[client]=true
			fmt.Println("totle connection ppol:- ", len(poll.Clients))
			for k, _ := range poll.Clients{
				fmt.Println(k)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined"})
			}
		case client := <-poll.Unregister:
		case client := <-poll.Broadcast:
		}
	}
}