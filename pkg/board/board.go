package board

import (
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
)

const (
	initialPositionFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

	sizeOfBoard = 8

	initialPositionKingWhiteX = 4
	initialPositionKingWhiteY = 7
	initialPositionKingBlackX = 4
	initialPositionKingBlackY = 0

	initialPositionRookKingWhiteX = 7
	initialPositionRookKingWhiteY = 7
	initialPositionRookKingBlackX = 7
	initialPositionRookKingBlackY = 0

	initialPositionRookQueenWhiteX = 0
	initialPositionRookQueenWhiteY = 7
	initialPositionRookQueenBlackX = 0
	initialPositionRookQueenBlackY = 0

	castlingKingSideX  = 7
	castlingQueenSideX = 0
)

var kingMoves = [8]values.Move{
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
}

var knightMoves = [8]values.Move{
	{X: -2, Y: -1},
	{X: -1, Y: -2},
	{X: -2, Y: 1},
	{X: 1, Y: -2},
	{X: -1, Y: 2},
	{X: 2, Y: -1},
	{X: 1, Y: 2},
	{X: 2, Y: 1},
}

type Board struct {
	rows               [sizeOfBoard][sizeOfBoard]piece.PieceContract
	currentColor       int
	enPassant          *values.Coord
	isSimulatation     bool
	checkWhite         bool
	checkBlack         bool
	castlingKingWhite  bool
	castlingKingBlack  bool
	castlingQueenWhite bool
	castlingQueenBlack bool
}

func NewBoard() Board {
	return Board{
		currentColor:       values.White,
		castlingKingWhite:  true,
		castlingKingBlack:  true,
		castlingQueenWhite: true,
		castlingQueenBlack: true,
	}
}

func (b *Board) Start() {
	b.rows = Build(initialPositionFen).rows
}
