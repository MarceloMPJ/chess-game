package rook_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece/rook"
)

func TestRook_Show(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		r := rook.NewRook(values.White)

		expected := '♜'
		result := r.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		r := rook.NewRook(values.Black)

		expected := '♖'
		result := r.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestRook_ShowFEN(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		r := rook.NewRook(values.White)

		expected := 'R'
		result := r.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		r := rook.NewRook(values.Black)

		expected := 'r'
		result := r.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestRook_IsValidMove(t *testing.T) {
	type params struct {
		origin values.Coord
		dest   values.Coord
	}

	type context struct {
		context  string
		expected bool
		args     []params
	}

	tests := []context{
		{
			context:  "when the move is valid",
			expected: true,
			args: []params{
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 4, Y: 7},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 4, Y: 3},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 0, Y: 4},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 6, Y: 4},
				},
			},
		},
		{
			context:  "when the move is invalid",
			expected: false,
			args: []params{
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 3, Y: 7},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 3, Y: 3},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 0, Y: 7},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 3, Y: 5},
				},
			},
		},
	}

	for _, test := range tests {
		r := rook.NewRook(values.White)

		t.Run(test.context, func(t *testing.T) {
			for _, arg := range test.args {
				result := r.IsValidMove(arg.origin, arg.dest)

				checkResult(t, result, test.expected)
			}
		})
	}
}

func checkResult(t *testing.T, result, expected bool) {
	t.Helper()

	if result != expected {
		t.Errorf("result: %t, expected: %t", result, expected)
	}
}
