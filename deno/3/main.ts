import { lineStream } from "../lib/stdin.ts";

let sum = 0;
let enabled = true;

for await (const line of lineStream) {
  const matches = line.matchAll(/(do|don't)\(\)|mul\((\d+),(\d+)\)/g);

  for (const match of matches) {
    console.log([...match]);
    const [, command, v1, v2] = match;

    if (command === "do") {
      enabled = true;
    } else if (command === "don't") {
      enabled = false;
    } else if (enabled) {
      sum += parseInt(v1 ?? "0") * parseInt(v2 ?? "0");
    }
  }
}

console.log("Sum of multiplications:", sum);
