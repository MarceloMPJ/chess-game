package bishop

type Bishop struct {
	color int
}

func NewBishop(color int) Bishop {
	return Bishop{color}
}

func (b *Bishop) Show() rune {
	if b.color == 0 {
		return '♝'
	}

	return '♗'
}
