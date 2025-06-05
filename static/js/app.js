// This file contains JavaScript code for client-side interactions, such as form validation and dynamic updates to the page.

document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("sudoku-form");
    const resultDiv = document.getElementById("result");

    form.addEventListener("submit", function(event) {
        event.preventDefault();
        const input = document.getElementById("sudoku-input").value;
        const board = parseInput(input);
        
        if (board) {
            fetch("/solve", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ board: board })
            })
            .then(response => response.json())
            .then(data => {
                if (data.solution) {
                    displaySolution(data.solution);
                } else {
                    resultDiv.innerHTML = "Nessuna soluzione trovata.";
                }
            })
            .catch(error => {
                console.error("Errore:", error);
                resultDiv.innerHTML = "Si è verificato un errore durante la risoluzione.";
            });
        } else {
            resultDiv.innerHTML = "Input non valido. Assicurati di inserire 81 numeri separati da virgola.";
        }
    });

    function parseInput(input) {
        const parts = input.split(",").map(num => parseInt(num.trim(), 10));
        if (parts.length !== 81 || parts.some(num => isNaN(num) || num < 0 || num > 9)) {
            return null;
        }
        return parts;
    }

    function displaySolution(solution) {
        let html = "<h3>Soluzione:</h3><table>";
        for (let i = 0; i < 9; i++) {
            html += "<tr>";
            for (let j = 0; j < 9; j++) {
                html += `<td>${solution[i * 9 + j]}</td>`;
            }
            html += "</tr>";
        }
        html += "</table>";
        resultDiv.innerHTML = html;
    }
});