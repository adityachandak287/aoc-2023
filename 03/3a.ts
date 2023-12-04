import { readFileSync } from "fs";

const input = readFileSync("03/input.txt", { encoding: "utf-8" });

const lines = input.split("\n");

let sum = 0;

for (let row = 0; row < lines.length; row++) {
  let currentNumber = 0;
  let nearSymbol = false;
  for (let col = 0; col < lines[0].length; col++) {
    const char = lines[row][col];
    if (char >= "0" && char <= "9") {
      currentNumber = currentNumber * 10 + Number(char);
      console.log("Found number", currentNumber, nearSymbol);
    } else {
      if (currentNumber !== 0 && nearSymbol) {
        sum += currentNumber;
        console.log("Adding number", currentNumber);
      }

      currentNumber = 0;
      nearSymbol = false;
    }

    if (currentNumber === 0) {
      continue;
    }

    for (
      let r = Math.max(0, row - 1);
      r <= Math.min(lines.length - 1, row + 1);
      r++
    ) {
      for (
        let c = Math.max(0, col - 1);
        c <= Math.min(lines[row].length - 1, col + 1);
        c++
      ) {
        if (r === row && c === col) {
          continue;
        }
        const sym = lines[r][c];
        console.log(currentNumber, sym, r, c);
        // if (sym !== "." && sym < "0" && sym > "9") {
        if (!(sym >= "0" && sym <= "9") && sym !== ".") {
          nearSymbol = true;
          console.log(
            "Found symbol near",
            currentNumber,
            r,
            c,
            "for",
            row,
            col
          );
        }
      }
    }
  }

  if (currentNumber !== 0 && nearSymbol) {
    sum += currentNumber;
  }

  currentNumber = 0;
  nearSymbol = false;
}

console.log("SUM", sum);
