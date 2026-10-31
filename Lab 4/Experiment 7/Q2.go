package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("negative number cannot have a real square root")
	}

	return math.Sqrt(n), nil
}

func parseNumber(input string) (float64, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}

	num, err := strconv.ParseFloat(input, 64)

	if err != nil || math.IsNaN(num) || math.IsInf(num, 0) {
		return 0, fmt.Errorf("invalid numeric input")
	}

	return num, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter numerator: ")
	input1, _ := reader.ReadString('\n')
	a, err := parseNumber(input1)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Print("Enter denominator: ")
		input2, _ := reader.ReadString('\n')
		b, err := parseNumber(input2)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			result, err := divide(a, b)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Division Result:", result)
			}
		}
	}

	fmt.Print("\nEnter number for square root: ")
	input3, _ := reader.ReadString('\n')
	n, err := parseNumber(input3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, err := squareRoot(n)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Square Root:", result)
		}
	}
}