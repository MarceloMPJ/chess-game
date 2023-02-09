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

func (r *Rook) ShowFEN() rune {
	if r.color == values.White {
		return 'R'
	}

	return 'r'
}

func (r *Rook) IsValidMove(origin, dest values.Coord) bool {
	return origin.X == dest.X || origin.Y == dest.Y
}
