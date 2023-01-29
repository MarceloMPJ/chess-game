package table

import (
	"github.com/MarceloMPJ/chess-game/src/piece"
	"github.com/MarceloMPJ/chess-game/src/piece/bishop"
	"github.com/MarceloMPJ/chess-game/src/piece/king"
	"github.com/MarceloMPJ/chess-game/src/piece/knight"
	"github.com/MarceloMPJ/chess-game/src/piece/pawn"
	"github.com/MarceloMPJ/chess-game/src/piece/queen"
	"github.com/MarceloMPJ/chess-game/src/piece/rook"
)

const (
	White = iota
	Black
)

type Table struct {
	rows [8][8]piece.PieceContract
}

func NewTable() Table {
	var rows [8][8]piece.PieceContract

	// Rooks
	rookWhiteL := rook.NewRook(White)
	rookWhiteR := rook.NewRook(White)
	rookBlackL := rook.NewRook(Black)
	rookBlackR := rook.NewRook(Black)

	// Bishops
	bishopWhiteL := bishop.NewBishop(White)
	bishopWhiteR := bishop.NewBishop(White)
	bishopBlackL := bishop.NewBishop(Black)
	bishopBlackR := bishop.NewBishop(Black)

	// Knights
	knightWhiteL := knight.NewKnight(White)
	knightWhiteR := knight.NewKnight(White)
	knightBlackL := knight.NewKnight(Black)
	knightBlackR := knight.NewKnight(Black)

	// Kings
	kingWhite := king.NewKing(White)
	kingBlack := king.NewKing(Black)

	// Queens
	queenWhite := queen.NewQueen(White)
	queenBlack := queen.NewQueen(Black)

	rows[0] = [8]piece.PieceContract{
		&rookBlackL,
		&knightBlackL,
		&bishopBlackL,
		&kingBlack,
		&queenBlack,
		&bishopBlackR,
		&knightBlackR,
		&rookBlackR,
	}

	rows[7] = [8]piece.PieceContract{
		&rookWhiteL,
		&knightWhiteL,
		&bishopWhiteL,
		&kingWhite,
		&queenWhite,
		&bishopWhiteR,
		&knightWhiteR,
		&rookWhiteR,
	}

	for i := 0; i < 8; i++ {
		pawnWhite := pawn.NewPawn(White)
		pawnBlack := pawn.NewPawn(Black)

		rows[6][i] = &pawnWhite
		rows[1][i] = &pawnBlack
	}

	return Table{rows: rows}
}

func (t *Table) Debug() (result string) {
	for i := 0; i < 8; i++ {
		line := ""

		for j := 0; j < 8; j++ {
			if t.rows[i][j] != nil {
				line += string(t.rows[i][j].Show())
			} else {
				line += " "
			}

			if j != 7 {
				line += " "
			}
		}
		result += line + "\n"
	}

	return
}
