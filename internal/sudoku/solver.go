package sudoku

import "fmt"

const N = 9

func SolveSudoku(board [][]int) bool {

	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(board, row, col, num) {
						board[row][col] = num
						if SolveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}

				return false
			}
		}
	}

	// Log The board is completely filled
	fmt.Println("Sudoku solved successfully!")
	// Log board
	for _, row := range board {
		fmt.Println(row)
	} // Return true if the Sudoku is solved
	return true

}

func isSafe(board [][]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if board[row][x] == num || board[x][col] == num {
			return false
		}
	}
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}
