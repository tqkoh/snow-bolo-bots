package bots

import (
	"github.com/gofrs/uuid"
)

type BotRandom0 struct {
	Id uuid.UUID
}

func (b *BotRandom0) SetId(id uuid.UUID) {
	b.Id = id
}

func (b *BotRandom0) GetName(_ int) string {
	return "BotRandom0"
}

func (b *BotRandom0) GetInput() Input {
	return Input{
		W:     true,
		A:     false,
		S:     false,
		D:     false,
		Left:  false,
		Right: false,
		Dy:    0,
		Dx:    0,
	}
}
