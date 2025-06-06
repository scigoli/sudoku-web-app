package web

import (
	"encoding/json"
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
		// Per richieste JSON
		w.Header().Set("Content-Type", "application/json")

		var req struct {
			Board [][]int `json:"board"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
			return
		}

		board := req.Board
		if !sudoku.SolveSudoku(board) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "No solution found"})
			return
		}

		json.NewEncoder(w).Encode(map[string]any{"solution": board})
		return
	}

	// Per GET, mostra la pagina HTML
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
	// Log the parsed board for debugging
	fmt.Println("Parsed Sudoku board:")
	for _, row := range board {
		fmt.Println(row)
	}
	return board, nil
}
