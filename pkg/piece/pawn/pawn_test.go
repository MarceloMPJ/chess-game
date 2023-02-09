package pawn_test

import (
	"testing"

	"github.com/MarceloMPJ/chess-game/libs/values"
	"github.com/MarceloMPJ/chess-game/pkg/piece/pawn"
)

func TestPawn_Show(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		p := pawn.NewPawn(values.White)

		expected := '♟'
		result := p.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		p := pawn.NewPawn(values.Black)

		expected := '♙'
		result := p.Show()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestPawn_ShowFEN(t *testing.T) {
	t.Run("when color is white", func(t *testing.T) {
		p := pawn.NewPawn(values.White)

		expected := 'P'
		result := p.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})

	t.Run("when color is black", func(t *testing.T) {
		p := pawn.NewPawn(values.Black)

		expected := 'p'
		result := p.ShowFEN()

		if result != expected {
			t.Errorf("result: %c, expected: %c", result, expected)
		}
	})
}

func TestPawn_IsValidMove(t *testing.T) {
	type params struct {
		origin values.Coord
		dest   values.Coord
	}

	type context struct {
		context  string
		expected bool
		args     []params
	}

	type contextColor struct {
		contextColor string
		color        int
		contexts     []context
	}

	tests := []contextColor{
		{
			contextColor: "when color is white",
			color:        values.White,
			contexts: []context{
				{
					context:  "when the move is valid",
					expected: true,
					args: []params{
						{
							values.Coord{X: 7, Y: 7},
							values.Coord{X: 7, Y: 6},
						},
						{
							values.Coord{X: 7, Y: 7},
							values.Coord{X: 7, Y: 5},
						},
					},
				},
				{
					context:  "when the move is invalid",
					expected: false,
					args: []params{
						{
							values.Coord{X: 7, Y: 7},
							values.Coord{X: 6, Y: 6},
						},
						{
							values.Coord{X: 7, Y: 7},
							values.Coord{X: 7, Y: 4},
						},
						{
							values.Coord{X: 6, Y: 6},
							values.Coord{X: 6, Y: 7},
						},
					},
				},
			},
		},
		{
			contextColor: "when color is black",
			color:        values.Black,
			contexts: []context{
				{
					context:  "when the move is valid",
					expected: true,
					args: []params{
						{
							values.Coord{X: 7, Y: 6},
							values.Coord{X: 7, Y: 7},
						},
						{
							values.Coord{X: 7, Y: 5},
							values.Coord{X: 7, Y: 7},
						},
					},
				},
				{
					context:  "when the move is invalid",
					expected: false,
					args: []params{
						{
							values.Coord{X: 6, Y: 6},
							values.Coord{X: 7, Y: 7},
						},
						{
							values.Coord{X: 7, Y: 4},
							values.Coord{X: 7, Y: 7},
						},
						{
							values.Coord{X: 6, Y: 7},
							values.Coord{X: 6, Y: 6},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.contextColor, func(t *testing.T) {
			p := pawn.NewPawn(test.color)

			for _, testContext := range test.contexts {
				t.Run(testContext.context, func(t *testing.T) {
					for _, arg := range testContext.args {
						result := p.IsValidMove(arg.origin, arg.dest)

						checkResult(t, result, testContext.expected)
					}
				})
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
