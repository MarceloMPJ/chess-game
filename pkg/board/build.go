package board

import (
	"strings"

	"github.com/MarceloMPJ/chess-game/libs/basic"
	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece"
	"github.com/MarceloMPJ/chess-game/pkg/piece/bishop"
	"github.com/MarceloMPJ/chess-game/pkg/piece/king"
	"github.com/MarceloMPJ/chess-game/pkg/piece/knight"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
	"github.com/MarceloMPJ/chess-game/pkg/piece/queen"
	"github.com/MarceloMPJ/chess-game/pkg/piece/rook"
)

func Build(fen string) Board {
	fenRows := strings.Split(fen, "/")

	b := NewBoard()
	for i, fenRow := range fenRows {
		j := 0

		for _, fenPiece := range fenRow {
			if basic.IsNumber(fenPiece) {
				n := basic.RuneToInt(fenPiece)
				j += n

				continue
			}

			p := buildPiece(fenPiece)
			b.rows[i][j] = p
			j++
		}
	}

	return b
}

func buildPiece(fenPiece rune) piece.PieceContract {
	switch fenPiece {
	case 'p':
		pc := pawn.NewPawn(values.Black)
		return &pc
	case 'P':
		pc := pawn.NewPawn(values.White)
		return &pc
	case 'r':
		pc := rook.NewRook(values.Black)
		return &pc
	case 'R':
		pc := rook.NewRook(values.White)
		return &pc
	case 'n':
		pc := knight.NewKnight(values.Black)
		return &pc
	case 'N':
		pc := knight.NewKnight(values.White)
		return &pc
	case 'b':
		pc := bishop.NewBishop(values.Black)
		return &pc
	case 'B':
		pc := bishop.NewBishop(values.White)
		return &pc
	case 'q':
		pc := queen.NewQueen(values.Black)
		return &pc
	case 'Q':
		pc := queen.NewQueen(values.White)
		return &pc
	case 'k':
		pc := king.NewKing(values.Black)
		return &pc
	case 'K':
		pc := king.NewKing(values.White)
		return &pc
	}

	return nil
}
