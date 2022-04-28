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
    │   └── chat_client.html  //chat client implementation `javaScript` web Socket.
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


 ## Client:
 > to test the chat-server we are using    javascript web socket client:

 ```code
        
        function connect() {
            
            var room_id = document.getElementById("room_id").value;
            var user_email = document.getElementById("user_email").value;
            var url = "ws://localhost:8080/ws/chat-server?room_id=" + room_id + "&user_email=" + user_email;

            console.log(url);

            socket = new WebSocket(url);

            socket.onopen = function () {
                output.innerHTML += "<div class=\"row\"> <div class=\"col\"> <p> Connected </p> </div></div>";
            };

            socket.onmessage = function (e) {
                output.innerHTML += "<div class=\"row\"> <div class=\"col\">" +
                    prettifyJson(e.data, true) +
                    "</div></div>";
            };

            socket.onclose = () => {
                console.log('Web Socket Connection Closed');
            };
        }

        function disconnect() {
            if (socket != null) {
                socket.close();
            }

        }

        function sendMessage() {
            if (socket == null) {
                console.log('Web Socket Connection Closed');
            }
            var user_name= document.getElementById("user_name").value;
            var msg = document.getElementById("msgToServer").value;
            socket.send(String(user_name)+" : "+msg);
        }

 ``` 