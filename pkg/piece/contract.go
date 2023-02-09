package piece

import "github.com/MarceloMPJ/chess-game/libs/values"

type PieceContract interface {
	Show() rune
	ShowFEN() rune
	IsValidMove(origin, dest values.Coord) bool
}
