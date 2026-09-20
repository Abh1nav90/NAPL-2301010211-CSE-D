package mathutil


func Power(b, a int) int {
	if a == 0 {
		return 1
	}
	if a < 0 {
		return -1 
	}
	power := 1
	for i := 0; i < a; i++ {
		power *= b
	}
	return power
}