const raw = await Deno.readTextFile("../../inputs/1/full.txt");

const arr1: number[] = [];
const arr2: number[] = [];

for (const row of raw.split(/\n/)) {
  const vals = row.split(/\s+/);

  if (vals.length !== 2) {
    throw Error("Row parsing error");
  }

  arr1.push(parseInt(vals[0]));
  arr2.push(parseInt(vals[1]));
}

arr1.sort((a, b) => a - b);
arr2.sort((a, b) => a - b);

const diff = arr1.map((val1, i) => Math.abs(val1 - arr2[i]));
const sum = diff.reduce((arr, val) => arr + val, 0);

console.log({ arr1, arr2, diff });
console.log("Final Result:", sum);
