/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-12-11
	* This file creates a calculator 
	*/


package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
		reader := bufio.NewReader(os.Stdin)

		// greeting user 
		fmt.Println("Welcome to my calculator program.")
		fmt.Println("Which operation would you like to perform today?")
		fmt.Println("a. add\nb. subtract\nc. multiply\nd. divide\ne. absolute value\nf. round\ng. raise to an exponent\nh. square root")

		fmt.Print("What operation do you want to choose: ")
		opInput, _ := reader.ReadString('\n')
		op := strings.TrimSpace(opInput)

		// addition 
		if op == "a" {
			fmt.Print("Enter first number: ")
			n1Input, _ := reader.ReadString('\n')
			n1, _ := strconv.ParseFloat(strings.TrimSpace(n1Input), 64)

			fmt.Print("Enter second number: ")
			n2Input, _ := reader.ReadString('\n')
			n2, _ := strconv.ParseFloat(strings.TrimSpace(n2Input), 64)

			fmt.Printf("%.2f + %.2f = %.2f\n", n1, n2, n1+n2)

			// subtraction 
		} else if op == "b" {
			fmt.Print("Enter first number: ")
			n1Input, _ := reader.ReadString('\n')
			n1, _ := strconv.ParseFloat(strings.TrimSpace(n1Input), 64)

			fmt.Print("Enter second number: ")
			n2Input, _ := reader.ReadString('\n')
			n2, _ := strconv.ParseFloat(strings.TrimSpace(n2Input), 64)

			fmt.Printf("%.2f - %.2f = %.2f\n", n1, n2, n1-n2)

			// multiplication
		} else if op == "c" {
			fmt.Print("Enter first number: ")
			n1Input, _ := reader.ReadString('\n')
			n1, _ := strconv.ParseFloat(strings.TrimSpace(n1Input), 64)

			fmt.Print("Enter second number: ")
			n2Input, _ := reader.ReadString('\n')
			n2, _ := strconv.ParseFloat(strings.TrimSpace(n2Input), 64)

			fmt.Printf("%.2f * %.2f = %.2f\n", n1, n2, n1*n2)

			// division
		} else if op == "d" {
			fmt.Print("Enter first number: ")
			n1Input, _ := reader.ReadString('\n')
			n1, _ := strconv.ParseFloat(strings.TrimSpace(n1Input), 64)

			fmt.Print("Enter second number (non-zero): ")
			n2Input, _ := reader.ReadString('\n')
			n2, _ := strconv.ParseFloat(strings.TrimSpace(n2Input), 64)

			if n2 == 0 {
				fmt.Println("Cannot divide by zero!")
			} else {
				fmt.Printf("%.2f / %.2f = %.2f\n", n1, n2, n1/n2)
			}

			// absolute value 
		} else if op == "e" {
			fmt.Print("Enter a number: ")
			nInput, _ := reader.ReadString('\n')
			n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

			fmt.Printf("Absolute value of %.2f is %.2f\n", n, math.Abs(n))

			// rounding 
		} else if op == "f" {
			fmt.Print("Enter a number: ")
			nInput, _ := reader.ReadString('\n')
			n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

			fmt.Printf("%.2f rounded is %.0f\n", n, math.Round(n))

			// raise to an exponent 
		} else if op == "g" {
			fmt.Print("Enter the base: ")
			baseInput, _ := reader.ReadString('\n')
			base, _ := strconv.ParseFloat(strings.TrimSpace(baseInput), 64)

			fmt.Print("Enter the exponent: ")
			expInput, _ := reader.ReadString('\n')
			exp, _ := strconv.ParseFloat(strings.TrimSpace(expInput), 64)

			fmt.Printf("%.2f raised to the power of %.2f is %.2f\n", base, exp, math.Pow(base, exp))

			// square root 
		} else if op == "h" {
			fmt.Print("Enter a non-negative number: ")
			nInput, _ := reader.ReadString('\n')
			n, _ := strconv.ParseFloat(strings.TrimSpace(nInput), 64)

			if n < 0 {
				fmt.Println("Cannot take square root of negative number!")
			} else {
				fmt.Printf("The square root of %.2f is %.2f\n", n, math.Sqrt(n))
			}
		} else {
			fmt.Println("Invalid option selected.")
		}

		fmt.Println("\nDone.")
}