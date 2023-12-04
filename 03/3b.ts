import { readFileSync } from "fs";
import assert from "assert";

const input = readFileSync("03/input.txt", { encoding: "utf-8" });

const lines = input.split("\n");

const allMatches: { num: number; sym: string }[] = [];

for (let row = 0; row < lines.length; row++) {
  let currentNumber = 0;
  let currentSymbol: string = "";

  for (let col = 0; col < lines[0].length; col++) {
    const char = lines[row][col];
    if (char >= "0" && char <= "9") {
      currentNumber = currentNumber * 10 + Number(char);
      console.log("Found number", currentNumber, currentSymbol);
    } else {
      if (currentNumber !== 0 && currentSymbol) {
        allMatches.push({ num: currentNumber, sym: currentSymbol });
        console.log("Adding number", currentNumber);
      }

      currentNumber = 0;
      currentSymbol = "";
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
        if (sym === "*") {
          currentSymbol = `${sym}_${r}_${c}`;
          console.log(
            `Found symbol ${sym} [${r},${c}] near ${currentNumber} [${row},${col}]`
          );
        }
      }
    }
  }

  if (currentNumber !== 0 && currentSymbol) {
    allMatches.push({ num: currentNumber, sym: currentSymbol });
  }

  currentNumber = 0;
  currentSymbol = "";
}

let sumOfGears = 0;
let map = new Map<string, number[]>();
for (const match of allMatches) {
  if (!map.has(match.sym)) {
    map.set(match.sym, [match.num]);
  } else {
    map.get(match.sym)?.push(match.num);
  }
}

map.forEach((nums, key) => {
  assert(nums.length <= 2, `${nums} length cannot be more than 2`);
  if (nums.length === 2) {
    sumOfGears += nums[0] * nums[1];
  }
});

console.log("SUM", sumOfGears);
