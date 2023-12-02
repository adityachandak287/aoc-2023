import { readFileSync } from "fs";
import assert from "assert";

const VALID_DIGITS = [
  "one",
  "two",
  "three",
  "four",
  "five",
  "six",
  "seven",
  "eight",
  "nine",
];

const input = readFileSync("01/input-1a.txt", { encoding: "utf-8" });

const lines = input.split("\n");

let sum = 0;

const checkAllSubstringsForDigits = (
  input: string,
  start: number,
  end: number
): number | null => {
  if (end - start + 1 < 3) {
    return null;
  }

  const inputSub = input.substring(start, end + 1);

  for (let head = 0; head < inputSub.length; head++) {
    for (let tail = head + 1; tail <= inputSub.length; tail++) {
      const sub = inputSub.substring(head, tail);
      assert(sub.length > 0, "Substring should not be empty string");

      const foundIdx = VALID_DIGITS.indexOf(sub);

      if (foundIdx !== -1) {
        return foundIdx + 1;
      }
    }
  }

  return null;
};

for (const line of lines) {
  assert(line.length > 0, "input line cannot be empty");
  assert(line.trim().length > 0, "trimmed input line cannot be empty");

  let first = -1;
  let last = -1;

  let startIdx = 0;

  for (let idx = 0; idx <= line.length; idx++) {
    const char = line.charAt(idx);

    let currentDigit: number = -1;
    if (char >= "0" && char <= "9") {
      currentDigit = Number(char);
    } else {
      const check = checkAllSubstringsForDigits(line, startIdx, idx);

      if (check === null) {
        continue;
      }

      currentDigit = check;
      startIdx = idx + 1;
    }

    if (currentDigit !== -1) {
      if (first === -1) {
        first = currentDigit;
      }

      last = currentDigit;
    }
  }

  const value = first * 10 + last;
  assert(value > 0, `final value ${value} should not be negative`);
  assert(value > 10, `final value ${value} should be more than 10`);
  // console.log(line, "VALUE", value, first, last);

  sum += value;
}

console.log(sum);
