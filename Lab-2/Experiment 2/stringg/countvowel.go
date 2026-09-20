package stringg


func CountVowels(s string) int {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	count := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(vowels); j++ {
			if rune(s[i]) == vowels[j] {
				count++
			}
		}
	}
	return count
}