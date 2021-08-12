package main
import(
	"github.com/gorilla/websocket"
)
//client represent one chatting user
type client struct{
	socket *websocket.Conn
	send chan []byte
	romm *room
}

