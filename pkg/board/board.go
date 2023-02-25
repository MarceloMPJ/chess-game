package board

import (
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
)

const initialPositionFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

type Board struct {
	rows         [sizeOfBoard][sizeOfBoard]piece.PieceContract
	currentColor int
	enPassant    *values.Coord
}

func NewBoard() Board {
	return Board{currentColor: values.White}
}

func (b *Board) Start() {
	b.rows = Build(initialPositionFen).rows
}
