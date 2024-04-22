package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {

	n := 8
	board := make([]int, n)
	solutions := []string{}
	var solve func(int, int, []int, *[]string, string)

	solve = func(n, col int, board []int, solutions *[]string, currentSolution string) {
		if col == n {
			*solutions = append(*solutions, currentSolution)
		return
		}
		for row := 0; row < n; row++ {
			if isSafe(board, row, col) {
				board[col] = row
				newSolution := currentSolution + strconv.Itoa(row+1)
				solve(n, col+1, board, solutions, newSolution)
			}
		}
	}
	solve(n, 0, board, &solutions, "")

	solutionsStr := strings.Join(solutions, "\n")
	return solutionsStr
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-1 == row-col || board[i]+i == row+col {
			return false
		}
	}
	return true
}