package main

import (
	"time"

	"github.com/gorilla/websocket"
)

//client represents a chatting user.
type client struct {
	//socket is websocket for this client
	socket *websocket.Conn
	//send is channel which is send message from server.
	send chan *message
	//room is chatroom where this client join.
	room *room
	//userData keeps data about user
	userData map[string]interface{}
}

//read reads data from websocket using ReadMessage. The sent data will be sent to the room's forward channel immediately.
func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			if avatarURL,ok:=c.userData["avatar_url"];ok{
				msg.AvatarURL=avatarURL.(string)
			}
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
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
