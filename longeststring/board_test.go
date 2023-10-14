package longeststring

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testData := []struct {
		name   string
		puzzle Board
		x      int
		y      int
		score  int
	}{
		{
			name: "Empty 3x3 board",
			puzzle: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			x:     0,
			y:     0,
			score: 1,
		},
		{
			name: "Full 3x3 board",
			puzzle: [][]int{
				{0, 1, 2},
				{5, 4, 3},
				{6, 7, 8},
			},
			x:     0,
			y:     0,
			score: 9,
		},
		{
			name: "Full reverse 3x3 board",
			puzzle: [][]int{
				{8, 7, 6},
				{3, 4, 5},
				{2, 1, 0},
			},
			x:     2,
			y:     2,
			score: 9,
		},
		{
			name: "Simple 3x3 board",
			puzzle: [][]int{
				{6, 2, 3},
				{7, 1, 4},
				{8, 0, 0},
			},
			x:     2,
			y:     1,
			score: 5,
		},
		{
			name: "Simple 4x3 board",
			puzzle: [][]int{
				{0, 6, 1},
				{2, 3, 0},
				{1, 4, 5},
				{0, 0, 6},
			},
			x:     3,
			y:     0,
			score: 7,
		},
		{
			name: "Loopy 8x6 board",
			puzzle: [][]int{
				{0, 8, 7, 6, 5, 0},
				{10, 9, 0, 1, 4, 0},
				{11, 8, 7, 2, 3, 0},
				{12, 13, 6, 3, 20, 21},
				{13, 14, 5, 4, 19, 22},
				{16, 15, 16, 17, 18, 23},
				{17, 30, 0, 28, 25, 24},
				{0, 29, 28, 27, 26, 27},
			},
			x:     1,
			y:     2,
			score: 31,
		},
	}
	for _, tt := range testData {
		name := tt.name
		puzzle := tt.puzzle
		x := tt.x
		y := tt.y
		score := tt.score
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			X, Y, Score := Solve(puzzle)
			if X != x || Y != y || Score != score {
				t.Errorf("wanted: 'x=%d, y=%d, score=%d', got 'x=%d, y=%d, score=%d'", x, y, score, X, Y, Score)
			}
		})
	}
}
