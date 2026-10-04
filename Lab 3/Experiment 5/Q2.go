package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks int
}

func modifyvalue(p *int){
	*p =*p + 20
}

func modifydataStruct(s *Student){
	s.Age = s.Age + 1
	s.Marks = s.Marks + 6
	s.Name = "Akhil"
}

func main() {
	// (1) Referencing and Dereferencing
	x := 20
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x: ", &x)
	fmt.Println("Value using pointer: ", *(&x))

	// Create pointer
	p := &x

	fmt.Println("pointer p", p)
	fmt.Println("Value using *p: ", *p)

	// (2) Pass pointer to function
	fmt.Println("Before modification: ", x)
	modifyvalue(&x)
	fmt.Println("After modification: ", x)

	// (3) Allocate struct using new()
	s := new(Student)
	fmt.Println("Student Details:")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Marks:", s.Marks)

	//after modification
	modifydataStruct(s)
	fmt.Println("After modification:")
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Marks:", s.Marks)

}
