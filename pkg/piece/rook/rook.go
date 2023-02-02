package rook

import "github.com/MarceloMPJ/chess-game/libs/values"

type Rook struct {
	color int
}

func NewRook(color int) Rook {
	return Rook{color}
}

func (r *Rook) Show() rune {
	if r.color == values.White {
		return '♜'
	}

	return '♖'
}
