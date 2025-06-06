package sudoku

import (
	"reflect"
	"testing"
)

func TestSolveSudoku_ValidPuzzle(t *testing.T) {
	input := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	board := make([][]int, 9)
	for i := range input {
		board[i] = make([]int, 9)
		copy(board[i], input[i])
	}
	t.Logf("Input board: %v", input)
	t.Log("Calling SolveSudoku...")
	ok := SolveSudoku(board)
	t.Logf("SolveSudoku returned: %v", ok)
	t.Logf("Solved board: %v", board)
	if !ok {
		t.Fatal("SolveSudoku failed to solve a valid puzzle")
	}
	if !reflect.DeepEqual(board, expected) {
		t.Errorf("SolveSudoku result mismatch.\nGot: %v\nExpected: %v", board, expected)
	} else {
		t.Log("Solution matches expected result.")
	}
}

func TestSolveSudoku_InvalidPuzzle(t *testing.T) {
	input := [][]int{
		{5, 3, 5, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	board := make([][]int, 9)
	for i := range input {
		board[i] = make([]int, 9)
		copy(board[i], input[i])
	}
	t.Logf("Input board (invalid): %v", input)
	t.Log("Calling SolveSudoku...")
	ok := SolveSudoku(board)
	t.Logf("SolveSudoku returned: %v", ok)
	t.Logf("Board after attempt: %v", board)
	if ok {
		t.Error("SolveSudoku should fail on an invalid puzzle")
	} else {
		t.Log("Correctly detected unsolvable puzzle.")
	}
}
