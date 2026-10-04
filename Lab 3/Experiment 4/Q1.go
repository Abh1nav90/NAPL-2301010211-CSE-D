package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) read() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	fmt.Scan(&p.Name)
	reader.ReadLine()

	fmt.Print("Enter Age: ")
	fmt.Scan(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scan(&p.Job)
	reader.ReadLine()

	fmt.Print("Enter Salary: ")
	fmt.Scan(&p.Salary)
}

func (p *Person) display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Println("Salary:", p.Salary)
}

func main() {

	var p1 Person
	fmt.Println("Enter details for person 1:")
	p1.read()
	fmt.Println("Details of person 1:")
	p1.display()

	var p2 Person
	fmt.Println("Enter details for person 2:")
	p2.read()
	fmt.Println("Details of person 2:")
	p2.display()
}
