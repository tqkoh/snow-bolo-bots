package main

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	println(os.Getenv("WS_URL"))
	c, _, err := websocket.DefaultDialer.Dial(os.Getenv("WS_URL"), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			println(t.String())
			// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	return
			// }
		}
	}
}
