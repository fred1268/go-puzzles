package longestpath

type Board [][]int

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func calculateScore(board, scores Board, x, y int, value int) int {
	if scores[x][y] != 0 {
		return scores[x][y]
	}
	score := 0
	value++
	if x > 0 && board[x-1][y] == value {
		score = max(score, calculateScore(board, scores, x-1, y, value))
	}
	if x < len(board)-1 && board[x+1][y] == value {
		score = max(score, calculateScore(board, scores, x+1, y, value))
	}
	if y > 0 && board[x][y-1] == value {
		score = max(score, calculateScore(board, scores, x, y-1, value))
	}
	if y < len(board[0])-1 && board[x][y+1] == value {
		score = max(score, calculateScore(board, scores, x, y+1, value))
	}
	score++
	scores[x][y] = score
	return score
}

func Solve(board Board) (X, Y, Score int) {
	MAX_X := len(board)
	MAX_Y := len(board[0])
	scores := make([][]int, MAX_X)
	for x := 0; x < MAX_X; x++ {
		scores[x] = make([]int, MAX_Y)
	}
	Score = -1
	for x := 0; x < MAX_X; x++ {
		for y := 0; y < MAX_Y; y++ {
			score := calculateScore(board, scores, x, y, board[x][y])
			if score > Score {
				X = x
				Y = y
				Score = score
			}
		}
	}
	return
}
