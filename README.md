CPE 490 Final Project
Izy Engel

To learn more about Go, I decided to build a simple chat server. While each client is connected to the server, each message sent by a client is simultaneously echoed to all clients from the host server. There are also some fun easter eggs- certain phrases sent to the room trigger the omnipotent host to dole out some wisdom or a greeting.


The go application leverages gin (github.com/gin-gonic/gin) to manage the HTTP framework and melody (https://gopkg.in/olahol/melody.v1) to manage the websocket. I also consulted this tutorial to get up and running: https://gabrieltanner.org/blog/realtime-chat-go-websockets

Video Demo: https://drive.google.com/file/d/10cwUPIe-erQ7v9iLkGtsDkMkXXViEaNn/view?usp=sharing


I pledge my honor that I have abided by the Stevens Honor System.
