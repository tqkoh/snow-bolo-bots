package bots

import "github.com/gofrs/uuid"

type Input struct {
	W     bool `json:"w"`
	A     bool `json:"a"`
	S     bool `json:"s"`
	D     bool `json:"d"`
	Left  bool `json:"left"`
	Right bool `json:"right"`
	Dy    int  `json:"dy"`
	Dx    int  `json:"dx"`
}

type Bot interface {
	SetId(id uuid.UUID)
	GetName(index int) string
	GetInput() Input
}

var Bots = []Bot{
	&BotRandom0{},
}
