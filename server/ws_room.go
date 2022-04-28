package server

import (
	"fmt"

	"github.com/google/uuid"
)

type RoomType string

const (
	RoomType_UNKOWN  RoomType = "UNKOWN"
	RoomType_DIRECT  RoomType = "DIRECT"
	RoomType_CHANNEL RoomType = "CHANNEL"
)

func (r RoomType) IsValid() error {
	if r != RoomType_CHANNEL && r != RoomType_DIRECT {
		return fmt.Errorf("room type UNKOWN")
	}
	return nil
}

// Room is a middleman between the Client and the hub.
type Room struct {
	ID uuid.UUID

	hub *Hub

	roomID string

	RoomType RoomType

	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan string

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newRoom(hub *Hub, roomID string) *Room {
	return &Room{
		hub:        hub,
		roomID:     roomID,
		broadcast:  make(chan string),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (r *Room) run() {
	defer func() {
		r.hub.unregister <- r
	}()

	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
			}
		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}

}
