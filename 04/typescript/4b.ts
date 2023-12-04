import { readFileSync } from "fs";

const data = readFileSync("04/input-04.txt", { encoding: "utf-8" });

const lines = data.split("\n");

const lineRegex = new RegExp(/:\s+| \| /);
const spaceRegex = new RegExp(/\s+/);

const cardResults: { cardId: string; matches: number }[] = [];
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

  cardResults.push({ cardId: cardId, matches: matches });
}

const cardCount = new Map<string, number>();

let resultIdx = 0;
for (const result of cardResults) {
  const numCards = 1 + (cardCount.get(result.cardId) ?? 0);

  for (let idx = resultIdx + 1; idx < resultIdx + 1 + result.matches; idx++) {
    const n = cardCount.get(cardResults[idx].cardId) ?? 0;
    cardCount.set(cardResults[idx].cardId, n + numCards);
  }

  cardCount.set(result.cardId, numCards);

  resultIdx++;
}

let score = 0;

for (const count of cardCount.values()) {
  score += count;
}

console.log("SCORE", score);
