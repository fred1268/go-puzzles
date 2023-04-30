package sudoku

import "fmt"

type Board [9][9]int

func display(board *Board) string {
	var result string
	for x := 0; x < 9; x++ {
		result = fmt.Sprintf("%s\n%d %d %d %d %d %d %d %d %d", result,
			board[x][0], board[x][1], board[x][2],
			board[x][3], board[x][4], board[x][5],
			board[x][6], board[x][7], board[x][8])
	}
	return result
}

func isPlayable(board *Board, x, y, value int) bool {
	for i := 0; i < 9; i++ {
		if board[i][y] == value {
			return false
		}
		if board[x][i] == value {
			return false
		}
	}
	x0, y0 := x/3, y/3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[x0*3+i][y0*3+j] == value {
				return false
			}
		}
	}
	return true
}

func Solve(board *Board) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == 0 {
				for value := 1; value <= 9; value++ {
					if isPlayable(board, x, y, value) {
						board[x][y] = value
						if !Solve(board) {
							board[x][y] = 0
						}
					}
				}
				if board[x][y] == 0 {
					return false
				}
			}
		}
	}
	return true
}
