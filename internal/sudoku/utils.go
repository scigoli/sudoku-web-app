package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

//const N = 9

func ParseInput(input string) ([][]int, error) {
	parts := strings.Split(input, ",")
	if len(parts) != N*N {
		return nil, fmt.Errorf("input must contain %d numbers separated by commas", N*N)
	}
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}
	for i, s := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || num < 0 || num > 9 {
			return nil, fmt.Errorf("invalid value: %s", s)
		}
		board[i/N][i%N] = num
	}
	return board, nil
}

func PrintBoard(board [][]int) string {
	var sb strings.Builder
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sb.WriteString(fmt.Sprintf("%d ", board[i][j]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
