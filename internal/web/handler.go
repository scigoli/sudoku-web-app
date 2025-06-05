package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"sudoku-web-app/internal/sudoku"
)

type SudokuHandler struct {
	Template *template.Template
}

func NewSudokuHandler() *SudokuHandler {
	tmpl := template.Must(template.ParseFiles("internal/web/templates/index.html"))
	return &SudokuHandler{Template: tmpl}
}

func (h *SudokuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		input := r.FormValue("sudokuInput")
		board, err := parseInput(input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if sudoku.SolveSudoku(board) {
			h.Template.Execute(w, board)
		} else {
			http.Error(w, "No solution found", http.StatusBadRequest)
		}
		return
	}

	h.Template.Execute(w, nil)
}

func parseInput(input string) ([][]int, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 81 {
		return nil, fmt.Errorf("input must contain 81 numbers separated by commas")
	}
	board := make([][]int, 9)
	for i := range board {
		board[i] = make([]int, 9)
	}
	for i, s := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || num < 0 || num > 9 {
			return nil, fmt.Errorf("invalid value: %s", s)
		}
		board[i/9][i%9] = num
	}
	return board, nil
}
