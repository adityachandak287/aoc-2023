import { readFileSync } from "fs";

const data = readFileSync("04/input-04.txt", { encoding: "utf-8" });

const lines = data.split("\n");

const lineRegex = new RegExp(/:\s+| \| /);
const spaceRegex = new RegExp(/\s+/);

let score = 0;
for (const line of lines) {
  const [cardId, winnerStr, numberStr] = line.split(lineRegex);
  const winners = winnerStr.split(spaceRegex).map(Number);
  const numbers = numberStr.split(spaceRegex).map(Number);

  let matches = 0;
  for (const num of numbers) {
    if (winners.includes(num)) {
      matches++;
    }
  }

  score += matches > 0 ? Math.pow(2, matches - 1) : 0;
}

console.log("SCORE", score);
