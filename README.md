# Sudoku Web Application

This project is a web application that allows users to input and solve Sudoku puzzles through a graphical interface. It utilizes Go for the backend and serves HTML, CSS, and JavaScript for the frontend.

## Project Structure

```
sudoku-web-app
├── cmd
│   └── server
│       └── main.go          # Entry point of the web application
├── internal
│   ├── sudoku
│   │   ├── solver.go        # Logic for solving Sudoku puzzles
│   │   └── utils.go         # Utility functions for Sudoku
│   └── web
│       ├── handler.go       # HTTP handlers for the web application
│       └── templates
│           └── index.html   # HTML template for the main page
├── static
│   ├── css
│   │   └── style.css        # CSS styles for the web application
│   └── js
│       └── app.js           # JavaScript for client-side interactions
├── go.mod                   # Go module definition
├── go.sum                   # Checksums for module dependencies
└── README.md                # Documentation for the project
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd sudoku-web-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/server/main.go
   ```

4. **Access the application:**
   Open your web browser and navigate to `http://localhost:8080`.

## Usage

- Input a Sudoku puzzle in the provided format (use `0` for empty cells).
- Click the "Solve" button to find the solution.
- The solution will be displayed on the same page.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.