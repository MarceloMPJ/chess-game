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
