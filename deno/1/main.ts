import { getInputStream } from "../lib/stdin.ts";

const arr1: number[] = [];
const arr2: number[] = [];

for await (const line of getInputStream()) {
  const vals = line.split(/\s+/);

  if (vals.length !== 2) {
    throw Error("Row parsing error");
  }

  arr1.push(parseInt(vals[0]));
  arr2.push(parseInt(vals[1]));
}

arr1.sort((a, b) => a - b);
arr2.sort((a, b) => a - b);

let distance = 0, similarity = 0;

for (let i = 0; i < arr1.length; i++) {
  distance += Math.abs(arr1[i] - arr2[i]);
  similarity += arr1[i] * arr2.filter((val) => val === arr1[i]).length;
}

console.log("Total distance:", distance);
console.log("Similarity score:", similarity);
