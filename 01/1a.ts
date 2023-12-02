import { readFileSync } from "fs";

const input = readFileSync("input-1a.txt", { encoding: "utf-8" });

const lines = input.split("\n");

let sum = 0;

for (const line of lines) {
  let first = -1;
  let last = -1;
  for (const char of line.split("")) {
    if (char >= "0" && char <= "9") {
      if (first === -1) {
        first = Number(char);
      }

      last = Number(char);
    }
    // console.log(first, last, first * 10 + last, line);
  }

  const value = first * 10 + last;

  sum += value;
}

console.log(sum);
