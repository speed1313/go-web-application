package main

import (
	"github.com/gorilla/websocket"
)

//client represents a chatting user.
type client struct {
	//socket is websocket for this client
	socket *websocket.Conn
	//send is channel which is send message from server.
	send chan []byte
	//room is chatroom where this client join.
	room *room
}

//read reads data from websocket using ReadMessage. The sent data will be sent to the room's forward channel immediately.
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

//write writes message which are sent by send channel to WebSocket to web browser by using WriteMessage.
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
