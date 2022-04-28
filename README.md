# chat-test



## Apis :

<!-- TABLE OF CONTENTS -->
## Table of Contents

- [Overview:](#About-The-Project)
- [Project Architecture](#project-architecture)
- [Folder Structure](#folder-structure)



## About The Project
`CHAT-TEST` is a pure [Golang][Golang] project, that manage a dead simple "chat" system. 

* `Server`: Receiving messages from a network interface (WebSocket) and forwarding them to all the connected clients
* `Client`: A process reading a string on `STDIN` and forwarding it to the server, and also receiving messages from the same server and writing them to `STDOUT`.


### Built With

* [Golang][Golang]
* [REST][REST]

<!-- PROJECT ARCHITECTURE -->
## Project Architecture

### Folder Structure

```
.
├── go.mod
├── go.sum
├── internal
│   └── path.go
├── main.go
├── README.md
├── vendor
└── server
    ├── handlers.go
    ├── response.go
    ├── server.go
    ├── view
    │   └── test.html
    ├── ws_client.go
    ├── ws_hub.go
    └── ws_room.go

```

## Entities

> This section will describe the entities of `CHAT-TEST` service

### Hub :
 register  set of chat rooms

 ### Room :
 register  set of the connected clients.

 ### Client:

 simple websocket  connection.


 [Test WebSocket Servers](https://www.piesocket.com/websocket-tester)