package king

import "github.com/MarceloMPJ/chess-game/libs/values"

type King struct {
	color int
}

func NewKing(color int) King {
	return King{color}
}

func (k *King) Show() rune {
	if k.color == values.White {
		return '♚'
	}

	return '♔'
}

func (k *King) IsValidMove(origin, dest values.Coord) bool {
	return (origin.X-dest.X <= 1 || dest.X-origin.X <= 1) &&
		(origin.Y-dest.Y <= 1 || dest.Y-origin.Y <= 1)
}
