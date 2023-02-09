package bishop_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece/bishop"
)

func TestBishop_Show(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		b := bishop.NewBishop(values.White)

		expected := '♝'
		result := b.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		b := bishop.NewBishop(values.Black)

		expected := '♗'
		result := b.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestBishop_ShowFEN(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		b := bishop.NewBishop(values.White)

		expected := 'B'
		result := b.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		b := bishop.NewBishop(values.Black)

		expected := 'b'
		result := b.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestBishop_IsValidMove(t *testing.T) {
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
					values.Coord{X: 7, Y: 7},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 0, Y: 0},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 7, Y: 1},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 1, Y: 7},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 5, Y: 5},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 5, Y: 3},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 3, Y: 3},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 2, Y: 6},
				},
			},
		},
		{
			context:  "when the move is invalid",
			expected: false,
			args: []params{
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 4, Y: 6},
				},
				{
					values.Coord{X: 3, Y: 4},
					values.Coord{X: 5, Y: 5},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 0, Y: 3},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 6, Y: 5},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 3, Y: 2},
				},
				{
					values.Coord{X: 4, Y: 4},
					values.Coord{X: 7, Y: 0},
				},
			},
		},
	}

	for _, test := range tests {
		b := bishop.NewBishop(values.White)

		t.Run(test.context, func(t *testing.T) {
			for _, arg := range test.args {
				result := b.IsValidMove(arg.origin, arg.dest)

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
