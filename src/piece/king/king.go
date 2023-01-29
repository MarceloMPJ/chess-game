package king

type King struct {
	color int
}

func NewKing(color int) King {
	return King{color}
}

func (k *King) Show() rune {
	if k.color == 0 {
		return '♚'
	}

	return '♔'
}
