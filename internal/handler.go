package internal

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/tqkoh/snow-bolo-bots/bots"
)

type Payload struct {
	Method string                 `json:"method,omitempty"`
	Args   map[string]interface{} `json:"args,omitempty"`
}

func Handle(bot bots.Bot, data []byte) {
	var payload Payload
	err := json.Unmarshal(data, &payload)
	if err != nil {
		log.Println(err)
		return
	}

	switch payload.Method {
	case "joinAccepted":
		log.Println("joinAccepted")
		bot.SetId(uuid.FromStringOrNil(payload.Args["id"].(string)))
	case "update":
	default:
	}
}
