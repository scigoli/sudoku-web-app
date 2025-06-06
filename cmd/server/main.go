package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"sudoku-web-app/internal/sudoku"

	"golang.org/x/time/rate"

	"github.com/gorilla/mux"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 10) // 10 richieste/secondo

func main() {
	r := mux.NewRouter()

	// Middleware per security headers
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("Content-Security-Policy", "default-src 'self'")
			next.ServeHTTP(w, r)
		})
	})

	r.Use(rateLimitMiddleware) // Applica il middleware di limitazione delle richieste

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
	fmt.Println("Received request to solve Sudoku")
	w.Header().Set("Content-Type", "application/json")

	// Verifica Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Content-Type must be application/json"})
		return
	}

	// Limita la dimensione del body
	r.Body = http.MaxBytesReader(w, r.Body, 1024) // max 1KB

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
	// Non esporre dettagli degli errori interni
	if err != nil {
		log.Printf("Error: %v", err)                                           // log interno
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // messaggio generico all'utente
		return
	}

	board := req.Board
	if err := validateSudokuBoard(board); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

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

func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func validateSudokuBoard(board [][]int) error {
	if len(board) != 9 {
		return fmt.Errorf("invalid board size")
	}
	for i, row := range board {
		if len(row) != 9 {
			return fmt.Errorf("invalid row size at %d", i)
		}
		for j, val := range row {
			if val < 0 || val > 9 {
				return fmt.Errorf("invalid value at position [%d,%d]", i, j)
			}
		}
	}
	return nil
}

/*
func solveSudoku(board [][]int) bool {
	// Implementa qui la logica per risolvere il Sudoku
	return false
}
*/
