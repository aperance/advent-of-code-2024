import { lineStream } from "../lib/stdin.ts";

const input = [];
for await (const line of lineStream) {
  const chars = line.split("");
  input.push(chars);
}

let wordCount = 0;
let crossCount = 0;
for (let x = 0; x < input.length; x++) {
  for (let y = 0; y < input[x].length; y++) {
    /** Check for XMAS */
    if (input[x][y] === "X") {
      for (let xDirection = -1; xDirection <= 1; xDirection++) {
        for (let yDirection = -1; yDirection <= 1; yDirection++) {
          let word = "";
          for (let magnitude = 0; magnitude <= 3; magnitude++) {
            const i = xDirection * magnitude + x;
            const j = yDirection * magnitude + y;
            word += input[i]?.[j];
          }
          if (word === "XMAS") {
            wordCount++;
          }
        }
      }
    }

    /** Check for X-MAS */
    if (input[x][y] === "A") {
      let slashCount = 0;
      for (let xVec = -1; xVec <= 1; xVec += 2) {
        for (let yVec = -1; yVec <= 1; yVec += 2) {
          if (
            input[x + xVec]?.[y + yVec] === "M" &&
            input[x + xVec * -1]?.[y + yVec * -1] === "S"
          ) {
            slashCount++;
          }
        }
      }
      if (slashCount === 2) {
        crossCount++;
      }
    }
  }
}

console.log("XMAS count:", wordCount);
console.log("X-MAS count:", crossCount);
