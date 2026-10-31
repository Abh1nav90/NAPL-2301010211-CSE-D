package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Square Goroutine Started")
	result := num * num
	ch <- fmt.Sprintf("Square of %d = %d", num, result)
	fmt.Println("Square Goroutine Completed")
}

func cube(num int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Cube Goroutine Started")
	result := num * num * num
	ch <- fmt.Sprintf("Cube of %d = %d", num, result)
	fmt.Println("Cube Goroutine Completed")
}

func fibonacci(n int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fibonacci Goroutine Started")

	a, b := 0, 1
	result := ""

	for i := 0; i < n; i++ {
		result += fmt.Sprintf("%d ", a)
		a, b = b, a+b
	}

	ch <- fmt.Sprintf("First %d Fibonacci numbers: %s", n, result)
	fmt.Println("Fibonacci Goroutine Completed")
}

func main() {
	ch := make(chan string, 3)

	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Main function started")

	go square(5, ch, &wg)
	go cube(3, ch, &wg)
	go fibonacci(7, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("Result:", result)
	}

	fmt.Println("Main function completed")
}