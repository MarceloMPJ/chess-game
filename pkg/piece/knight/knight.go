package knight

import "github.com/MarceloMPJ/chess-game/libs/values"

type Knight struct {
	color int
}

func NewKnight(color int) Knight {
	return Knight{color}
}

func (n *Knight) Show() rune {
	if n.color == values.White {
		return '♞'
	}

	return '♘'
}

func (n *Knight) IsValidMove(origin, dest values.Coord) bool {
	if (origin.X-2 == dest.X || origin.X+2 == dest.X) && (origin.Y-1 == dest.Y || origin.Y+1 == dest.Y) {
		return true
	}

	if (origin.Y-2 == dest.Y || origin.Y+2 == dest.Y) && (origin.X-1 == dest.X || origin.X+1 == dest.X) {
		return true
	}

	return false
}
