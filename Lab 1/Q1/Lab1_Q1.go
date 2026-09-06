package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Abhinav Jha")
	fmt.Println("Integer Operation")
	a := 18
	b := 6

	fmt.Printf("Addition: %d \n", a+b)
	fmt.Printf("Subtraction: %d \n", a-b)
	fmt.Printf("Multiplication: %d \n", a*b)
	fmt.Printf("Division: %d \n", a/b)

	fmt.Println("Float Operation")
	c := 12.5
	d := 2.5

	fmt.Printf("Addition: %f \n", c+d)
	fmt.Printf("Substraction: %f \n", c-d)
	fmt.Printf("Multiplication: %f \n", c*d)
	fmt.Printf("Division: %f \n", c/d)
}
