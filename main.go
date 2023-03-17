package main

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tqkoh/snow-bolo-bots/bots"
	"github.com/tqkoh/snow-bolo-bots/internal"
)

func main() {
	println(os.Getenv("WS_URL"))

	conns := make([]*websocket.Conn, len(bots.Bots))
	done := make(chan int)
	for i, bot := range bots.Bots {
		c, _, err := websocket.DefaultDialer.Dial(os.Getenv("WS_URL"), nil)
		if err != nil {
			log.Fatal("dial: ", err)
		}
		conns[i] = c

		c.WriteJSON(map[string]interface{}{
			"method": "join",
			"args": map[string]interface{}{
				"name": bot.GetName(i),
			},
		})
		if i > 0 {
			c.WriteJSON(map[string]interface{}{
				"method": "active",
				"args": map[string]interface{}{
					"active": false,
				},
			})
		}
		go func(c *websocket.Conn, bot bots.Bot, handle func(bots.Bot, []byte)) {
			defer func() { done <- 1 }()
			defer c.Close()
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("read: ", err)
					return
				}
				handle(bot, message)
			}
		}(c, bot, internal.Handle)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	doneBotNum := 0
	for {
		select {
		case <-done:
			doneBotNum++
			if doneBotNum >= len(bots.Bots) {
				return
			}
		case t := <-ticker.C:
			for i, bot := range bots.Bots {
				conns[i].WriteJSON(map[string]interface{}{
					"method": "input",
					"args":   bot.GetInput(),
				})
			}
			println(t.String())
			// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	return
			// }
		}
	}
}
