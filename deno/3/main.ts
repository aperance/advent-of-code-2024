import { getInput } from "../lib/stdin.ts";

const result = (await getInput()).matchAll(/mul\((\d+),(\d+)\)/g).reduce(
  (acc, match) => {
    console.log(match[0]);
    return acc + (parseInt(match[1]) * parseInt(match[2]));
  },
  0,
);

console.log("Sum of multiplications:", result);
