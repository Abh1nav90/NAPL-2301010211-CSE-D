package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getInt(prompt string) int {
	for {
		var input string
		fmt.Print(prompt)
		fmt.Scanln(&input)

		value, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid input! Enter an integer.")
			continue
		}

		if value < -1000 || value > 1000 {
			fmt.Println("Out of range! Enter a value between -1000 and 1000.")
			continue
		}

		return value
	}
}

func getFloat(prompt string) float64 {
	for {
		var input string
		fmt.Print(prompt)
		fmt.Scanln(&input)

		value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil {
			fmt.Println("Invalid input! Enter a number.")
			continue
		}

		if value < -1000 || value > 1000 {
			fmt.Println("Out of range! Enter a value between -1000 and 1000.")
			continue
		}

		return value
	}
}

func main() {
	for {
		fmt.Println("\n===== SIMPLE CALCULATOR =====")
		fmt.Println("1. Perform operations using Int")
		fmt.Println("2. Perform operations using Float")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("\n--- Integer Operations ---")

			a := getInt("Enter first integer: ")
			b := getInt("Enter second integer: ")

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

			if b != 0 {
				fmt.Println("Division:", a/b)
			} else {
				fmt.Println("Division: Cannot divide by zero")
			}

		} else if choice == 2 {
			fmt.Println("\n--- Floating-Point Operations ---")

			a := getFloat("Enter first number: ")
			b := getFloat("Enter second number: ")

			fmt.Println("Addition:", a+b)
			fmt.Println("Subtraction:", a-b)
			fmt.Println("Multiplication:", a*b)

			if b != 0 {
				fmt.Println("Division:", a/b)
			} else {
				fmt.Println("Division: Cannot divide by zero")
			}

		} else if choice == 3 {
			fmt.Println("Exiting program...")
			break

		} else {
			fmt.Println("Invalid choice! Please select 1, 2, or 3.")
		}
	}
}
