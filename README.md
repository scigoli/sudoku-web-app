# Sudoku Web Application

Questa applicazione web permette di inserire e risolvere puzzle Sudoku tramite un'interfaccia grafica moderna. Il backend è scritto in Go, mentre il frontend utilizza HTML, CSS e JavaScript.

## Funzionalità principali

- **Inserimento dati tramite griglia 9x9**: puoi inserire i numeri direttamente nelle celle della griglia. Le celle vuote sono considerate come zeri.
- **Pulsanti di controllo**:
  - **Solve**: risolve il Sudoku e mostra la soluzione in una tabella con evidenza dei valori di input.
  - **Test**: popola la griglia con un esempio di Sudoku predefinito.
  - **Clear**: svuota tutte le celle della griglia.
- **Visualizzazione soluzione**: la soluzione viene mostrata in una tabella con bordi spessi per i blocchi 3x3 e i valori di input evidenziati in grassetto e con sfondo azzurrino.
- **Tempo di risoluzione**: sotto la soluzione viene indicato il tempo impiegato per la risoluzione (in millisecondi).

## Project Structure

```
sudoku-web-app
├── cmd
│   └── server
│       └── main.go          # Entry point dell'applicazione web
├── internal
│   ├── sudoku
│   │   ├── solver.go        # Logica di risoluzione Sudoku
│   │   └── utils.go         # Utility
│   └── web
│       ├── handler.go       # HTTP handler
│       └── templates
│           └── index.html   # Template HTML principale
├── static
│   ├── css
│   │   └── style.css        # Stili CSS
│   └── js
│       └── app.js           # JavaScript per interazione client
├── go.mod                   # Modulo Go
├── go.sum                   # Checksum dipendenze
└── README.md                # Documentazione
```

## Setup Instructions

1. **Clona il repository:**
   ```
   git clone https://github.com/scigoli/sudoku-web-app
   cd sudoku-web-app
   ```

2. **Installa le dipendenze:**
   ```
   go mod tidy
   ```

3. **Avvia l'applicazione:**
   ```
   go run cmd/server/main.go
   ```

4. **Accedi all'applicazione:**
   Apri il browser e vai su `http://localhost:8080`.

## Utilizzo

- **Inserisci i numeri** nella griglia 9x9. Lascia vuote le celle che vuoi siano considerate come zeri.
- Puoi usare il pulsante **Test** per caricare un esempio di Sudoku.
- Premi **Solve** per vedere la soluzione e il tempo di risoluzione.
- Premi **Clear** per svuotare la griglia e inserire un nuovo puzzle.

## Schermate
![Input](https://github.com/scigoli/sudoku-web-app/blob/main/test-input.png?raw=true)
![Output](https://github.com/scigoli/sudoku-web-app/blob/main/test-output.png?raw=true)


## Unit test
   ```
   go test ./internal/sudoku
   ```

## Come contribuire

Contributi sono benvenuti! Apri una issue o una pull request per suggerimenti o correzioni.

## Licenza

Questo progetto è distribuito con licenza MIT. Vedi il file LICENSE per dettagli.
