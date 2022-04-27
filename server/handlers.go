package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (svc *HttpService) getChatServer(c *gin.Context) {

	userEmail := c.Query(`user_email`)
	if len(userEmail) == 0 {
		BindJsonErr(c, errors.New("userEmail should not be empty"))
		return
	}

	roomID := c.Query(`room_id`)
	if len(roomID) == 0 {
		BindJsonErr(c, errors.New("room_id should not be empty"))
		return
	}

	room := svc.hub.getRoomByEmailAndRoomID(userEmail, roomID)
	if room == nil {
		room = newRoom(svc.hub, userEmail, roomID)
		room.hub.register <- room
		// run room
		go room.run()
	}

	// now upgrade the HTTP  connection to a websocket connection
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		e := "could not upgrade to websocket connection"
		svc.Log.WithError(err).Error(e)
		AbortWithError(c, 500, e, err)
		return
	}

	client := &Client{ID: uuid.New(), room: room, conn: conn, send: make(chan string)}
	client.room.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	client.readPump()
}
