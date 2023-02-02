package bishop

import "github.com/MarceloMPJ/chess-game/libs/values"

type Bishop struct {
	color int
}

func NewBishop(color int) Bishop {
	return Bishop{color}
}

func (b *Bishop) Show() rune {
	if b.color == values.White {
		return '♝'
	}

	return '♗'
}
