package main

import (
	"example.com/app/mathutil"
	"example.com/app/stringg"
	"fmt"
)

func main() {
	fmt.Println("Enter a string to reverse:")
	var a_str string
	fmt.Scanln(&a_str)
	fmt.Println("Reversed string:", stringg.Reverse(a_str))

	fmt.Println("Enter a string to count vowels:")
	var countv string
	fmt.Scanln(&countv)
	fmt.Println("Number of vowels:", stringg.CountVowels(countv))

	fmt.Println("Enter a number to calculate factorial:")
	var fact int
	fmt.Scanln(&fact)
	fmt.Println("Factorial:", mathutil.Factorial(fact))

	fmt.Println("Enter the base and exponent to calculate power:")
	var base, exp int
	fmt.Scanln(&base, &exp)
	fmt.Println("Power:", mathutil.Power(base, exp))
}