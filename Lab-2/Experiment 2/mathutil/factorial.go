package mathutil

func Factorial(n int) int {
	if n < 0 {
		return -1 
	}
	if n == 0 || n == 1 {
		return 1
	}
	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}
	return factorial
}