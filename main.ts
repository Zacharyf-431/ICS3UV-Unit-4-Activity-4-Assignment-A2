/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-12-11
 * @fileoverview: This file creates a calculator 
 */

// greeting user and giving options 
console.log("Welcome to my calculator program.");
console.log("Which operation would you like to perform today?");
console.log("a. add\nb. subtract\nc. multiply\nd. divide\ne. absolute value\nf. round\ng. raise to an exponent\nh. square root");

let operation = prompt("What operation do you want to choose: ") || "";

// addition 
if (operation === "a") {
  let n1 = parseFloat(prompt("Enter first number: ") || "0");
  let n2 = parseFloat(prompt("Enter second number: ") || "0");
  console.log(`${n1} + ${n2} = ${n1 + n2}`);

  // sutraction 
} else if (operation === "b") {
  let n1 = parseFloat(prompt("Enter first number: ") || "0");
  let n2 = parseFloat(prompt("Enter second number: ") || "0");
  console.log(`${n1} - ${n2} = ${n1 - n2}`);

  // multiplication 
} else if (operation === "c") {
  let n1 = parseFloat(prompt("Enter first number: ") || "0");
  let n2 = parseFloat(prompt("Enter second number: ") || "0");
  console.log(`${n1} * ${n2} = ${n1 * n2}`);

  // division
} else if (operation === "d") {
  let n1 = parseFloat(prompt("Enter first number: ") || "0");
  let n2 = parseFloat(prompt("Enter second number (non-zero): ") || "1");
  if (n2 === 0) {
    console.log("Cannot divide by zero!");
  } else {
    console.log(`${n1} / ${n2} = ${n1 / n2}`);
  }

  // absolute value 
} else if (operation === "e") {
  let n = parseFloat(prompt("Enter a number: ") || "0");
  console.log(`Absolute value of ${n} is ${Math.abs(n)}`);

  // rounding 
} else if (operation === "f") {
  let n = parseFloat(prompt("Enter a number: ") || "0");
  console.log(`${n} rounded is ${Math.round(n)}`);

  // raise to an exponent 
} else if (operation === "g") {
  let base = parseFloat(prompt("Enter the base: ") || "0");
  let exp = parseFloat(prompt("Enter the exponent: ") || "0");
  console.log(`${base} raised to the power of ${exp} is ${Math.pow(base, exp)}`);

  // square root 
} else if (operation === "h") {
  let n = parseFloat(prompt("Enter a non-negative number: ") || "0");
  if (n < 0) {
    console.log("Cannot take square root of negative number!");
  } else {
    console.log(`The square root of ${n} is ${Math.sqrt(n)}`);
  }
} else {
  console.log("Invalid option selected.");
}

console.log("\nDone.");