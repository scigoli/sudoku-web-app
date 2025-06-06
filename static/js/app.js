// This file contains JavaScript code for client-side interactions, such as form validation and dynamic updates to the page.

document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("sudoku-form");
    const resultDiv = document.getElementById("result");

    form.addEventListener("submit", function(event) {
        event.preventDefault();
        const board = [];
        for (let i = 0; i < 9; i++) {
            const row = [];
            for (let j = 0; j < 9; j++) {
                const cell = document.getElementById(`cell-${i}-${j}`);
                let value = parseInt(cell.value, 10);
                if (isNaN(value) || value < 1 || value > 9) {
                    value = 0;
                }
                row.push(value);
            }
            board.push(row);
        }
        // Salva la board di input per evidenziare i valori originali
        window.lastInputBoard = board.map(row => row.slice());
        fetch("/solve", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ board: board })
        })
        .then(response => response.json())
        .then(data => {
            if (data.solution) {
                displaySolution(data.solution, window.lastInputBoard, data.elapsedMs);
            } else {
                resultDiv.innerHTML = "Nessuna soluzione trovata.";
            }
        })
        .catch(error => {
            console.error("Errore:", error);
            resultDiv.innerHTML = "Si è verificato un errore durante la risoluzione.";
        });
    });

    function displaySolution(solution, inputBoard, elapsedMs) {
        let html = "<h3>Soluzione:</h3><table class='sudoku-result-table'>";
        for (let i = 0; i < 9; i++) {
            html += "<tr>";
            for (let j = 0; j < 9; j++) {
                // Calcola le classi per i bordi spessi dei blocchi 3x3
                let classes = [];
                if (j % 3 === 0) classes.push("block-left");
                if (i % 3 === 0) classes.push("block-top");
                if (j === 8) classes.push("block-right");
                if (i === 8) classes.push("block-bottom");
                // Evidenzia i valori di input
                if (inputBoard && inputBoard[i][j] !== 0) {
                    classes.push("input-cell");
                }
                html += `<td class='${classes.join(" ")}' >${solution[i][j]}</td>`;
            }
            html += "</tr>";
        }
        html += "</table>";
        if (typeof elapsedMs !== "undefined") {
            html += `<div class='solution-time'>Tempo impiegato: <b>${elapsedMs}</b> ms</div>`;
        }
        resultDiv.innerHTML = html;
    }

    // Genera la griglia di input dinamicamente
    const tbody = document.getElementById("sudoku-tbody");
    if (tbody) {
        for (let i = 0; i < 9; i++) {
            const tr = document.createElement("tr");
            for (let j = 0; j < 9; j++) {
                const td = document.createElement("td");
                let classes = [];
                if (j % 3 === 0) classes.push("block-left");
                if (i % 3 === 0) classes.push("block-top");
                if (j === 8) classes.push("block-right");
                if (i === 8) classes.push("block-bottom");
                td.className = classes.join(" ");
                const input = document.createElement("input");
                input.type = "number";
                input.min = "1";
                input.max = "9";
                input.className = "sudoku-cell";
                input.id = `cell-${i}-${j}`;
                input.style = "width:2em;text-align:center;";
                td.appendChild(input);
                tr.appendChild(td);
            }
            tbody.appendChild(tr);
        }
    }

    // Pulsante Test
    document.getElementById("test-btn").addEventListener("click", function() {
        const testVals = "5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9".split(",");
        for (let i = 0; i < 9; i++) {
            for (let j = 0; j < 9; j++) {
                const idx = i * 9 + j;
                const val = testVals[idx].trim();
                const cell = document.getElementById(`cell-${i}-${j}`);
                cell.value = val !== "0" ? val : "";
            }
        }
    });

    // Pulsante Clear
    document.getElementById("clear-btn").addEventListener("click", function() {
        for (let i = 0; i < 9; i++) {
            for (let j = 0; j < 9; j++) {
                const cell = document.getElementById(`cell-${i}-${j}`);
                cell.value = "";
            }
        }
    });
});