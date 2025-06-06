package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"

	"sudoku-web-app/internal/sudoku"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/solve", solveHandler).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
	// Log the server start
	fmt.Println("Server started on :8080")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/web/templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		// Log the error for debugging
		fmt.Println("Error loading template:", err)
		return
	}
	// Log the request for debugging
	fmt.Println("Received request for home page")
	tmpl.Execute(w, nil)
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	// Log the request for debugging
	fmt.Println("Received request to solve Sudoku!")
	w.Header().Set("Content-Type", "application/json")

	start := time.Now() // tempo inizio

	// Read and log the raw body
	bodyBytes, _ := io.ReadAll(r.Body)
	fmt.Println("Raw body:", string(bodyBytes))
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Decodifica il JSON dal body
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
	fmt.Println("Sudoku board before solving:")
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println("Starting to solve Sudoku...")
	solved := sudoku.SolveSudoku(board)
	fmt.Println("Sudoku solved:", solved)

	elapsed := time.Since(start).Milliseconds()

	if !solved {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "No solution found"})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"solution": board, "elapsedMs": elapsed})
}

/*
func solveSudoku(board [][]int) bool {
	// Implementa qui la logica per risolvere il Sudoku
	return false
}
*/
