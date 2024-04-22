package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	n := 8
	board := make([]int, n)
	solutions := []string{}
	var solve func(int)

	solve = func(col int) {
		if col == n {
			var solution strings.Builder
			for _, row := range board {
				solution.WriteString(strconv.Itoa(row + 1))
			}
			solutions = append(solutions, solution.String())
			return
		}
		for row := 0; row < n; row++ {
			if isSafe(board, row, col) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)

	solutionsStr := strings.Join(solutions, "\n")
	return solutionsStr
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-i == row-col || board[i]+i == row+col {
			return false
		}
	}
	return true
}