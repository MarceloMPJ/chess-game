package board

import (
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
	"github.com/MarceloMPJ/chess-game/pkg/piece/bishop"
	"github.com/MarceloMPJ/chess-game/pkg/piece/king"
	"github.com/MarceloMPJ/chess-game/pkg/piece/knight"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
	"github.com/MarceloMPJ/chess-game/pkg/piece/queen"
	"github.com/MarceloMPJ/chess-game/pkg/piece/rook"
)

type Board struct {
	rows         [8][8]piece.PieceContract
	currentColor int
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) Start() {
	var rows [8][8]piece.PieceContract

	// Rooks
	rookWhiteL := rook.NewRook(values.White)
	rookWhiteR := rook.NewRook(values.White)
	rookBlackL := rook.NewRook(values.Black)
	rookBlackR := rook.NewRook(values.Black)

	// Bishops
	bishopWhiteL := bishop.NewBishop(values.White)
	bishopWhiteR := bishop.NewBishop(values.White)
	bishopBlackL := bishop.NewBishop(values.Black)
	bishopBlackR := bishop.NewBishop(values.Black)

	// Knights
	knightWhiteL := knight.NewKnight(values.White)
	knightWhiteR := knight.NewKnight(values.White)
	knightBlackL := knight.NewKnight(values.Black)
	knightBlackR := knight.NewKnight(values.Black)

	// Kings
	kingWhite := king.NewKing(values.White)
	kingBlack := king.NewKing(values.Black)

	// Queens
	queenWhite := queen.NewQueen(values.White)
	queenBlack := queen.NewQueen(values.Black)

	rows[0] = [8]piece.PieceContract{
		&rookBlackL,
		&knightBlackL,
		&bishopBlackL,
		&queenBlack,
		&kingBlack,
		&bishopBlackR,
		&knightBlackR,
		&rookBlackR,
	}

	rows[7] = [8]piece.PieceContract{
		&rookWhiteL,
		&knightWhiteL,
		&bishopWhiteL,
		&queenWhite,
		&kingWhite,
		&bishopWhiteR,
		&knightWhiteR,
		&rookWhiteR,
	}

	for i := 0; i < 8; i++ {
		pawnWhite := pawn.NewPawn(values.White)
		pawnBlack := pawn.NewPawn(values.Black)

		rows[6][i] = &pawnWhite
		rows[1][i] = &pawnBlack
	}

	b.currentColor = values.White
	b.rows = rows
}
