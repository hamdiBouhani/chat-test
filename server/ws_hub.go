package server

type Hub struct {
	// Registered clients.
	rooms map[*Room]bool //check this bool!!

	// Register requests from the Rooms.
	register chan *Room

	// Unregister requests from Rooms.
	unregister chan *Room
}

func newHub() *Hub {
	return &Hub{

		register:   make(chan *Room),
		unregister: make(chan *Room),
		rooms:      make(map[*Room]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case room := <-h.register:
			h.rooms[room] = true
		case room := <-h.unregister:
			if _, ok := h.rooms[room]; ok {
				delete(h.rooms, room)
			}
		}
	}
}

func (h *Hub) getRoomByRoomID(roomID string) *Room {
	for room := range h.rooms {
		if room.roomID == roomID {
			return room
		}
	}
	return nil
}

func (h *Hub) broadcastMessageInRoom(email string, roomID string, msg string) {
	room := h.getRoomByRoomID(roomID)
	if room != nil {
		room.broadcast <- msg
	}
}
