package main

import (
	static "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"fmt"
	"rsc.io/quote"                                                                                                                                                                                       
	"strings"
)


func main() {
	r := gin.Default()
	m := melody.New()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
 

	//messages format: {"username":"Iz","content":"hi"}
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		
		fmt.Println(string(msg))

		m.Broadcast(msg)

		stringMsg:= string(msg)
		stringMsg = strings.ToLower(stringMsg)
		if(strings.Contains(stringMsg, "hi") || strings.Contains(stringMsg, "hello") || strings.Contains(stringMsg, "hey") || strings.Contains(stringMsg, "howdy")) {
			welcomeMsg  := []byte("{\"username\":\"Host\",\"content\":\"" + quote.Hello() + "\"}")
			m.Broadcast(welcomeMsg)
		}
		 if(strings.Contains(stringMsg, "proverb") || strings.Contains(stringMsg, "advice") || strings.Contains(stringMsg, "wisdom")) {
			welcomeMsg  := []byte("{\"username\":\"Host\",\"content\":\"" + quote.Go() + "\"}")
			m.Broadcast(welcomeMsg)
		}
		 if(strings.Contains(stringMsg, "glass")) {
			welcomeMsg  := []byte("{\"username\":\"Host\",\"content\":\"" + quote.Glass() + "\"}")
			m.Broadcast(welcomeMsg)
		}
		 if(strings.Contains(stringMsg, "optimize") || strings.Contains(stringMsg, "optimization") ) {
			welcomeMsg  := []byte("{\"username\":\"Host\",\"content\":\"" + quote.Opt() + "\"}")
			m.Broadcast(welcomeMsg)
		}
		if(strings.Contains(stringMsg, "meaning of life?") ) {
			welcomeMsg  := []byte("{\"username\":\"Host\",\"content\":\"" + "42" + "\"}")
			m.Broadcast(welcomeMsg)
		}
		
	})
	
	r.Run(":5000")
}


