const fs = require("node:fs");

console.time();

const records = fs.readFileSync("./numbers.csv");

const numbers = records?.toString().split(",").map(Number);

const sum = numbers.reduce((prev, curr) => prev + curr, 0);

console.log(sum);
console.timeEnd();
