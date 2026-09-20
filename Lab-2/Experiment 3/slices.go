package main

import "fmt"

func main() {

	s := []string{"Abhinav", "Nikhil", "Lakshay", "Harshit", "Ayush"}

	fmt.Println("Slice of strings:", s)

	var name string
	fmt.Println("Enter a name to add in the slice:")
	fmt.Scanln(&name)

	// Insert
	s = append(s, name)
	fmt.Println("Updated slice of strings:", s)

	// Delete
	index := 1
	s = append(s[:index], s[index+1:]...)
	fmt.Println("After removing index", index, ":", s)

	// Update
	s[1] = "Rohan"
	fmt.Println("After updating:", s)

	// Map
	marks := map[string]int{
		"Math":    80,
		"Science": 85,
		"English": 90,
	}

	fmt.Println("Initial map:", marks)

	// Insert
	marks["Computer"] = 95
	fmt.Println("After inserting:", marks)

	// Delete
	delete(marks, "English")
	fmt.Println("After deleting English:", marks)

	// Lookup
	value, exists := marks["Math"]

	if exists {
		fmt.Println("Math marks:", value)
		fmt.Println("After lookup:", marks)
	} else {
		fmt.Println("Math not found")
	}

}
