package EX05

import (
	"fmt"
)

// MultValuesBetween...
func MultValuesBetween(x, y int) [3][]int {
	var initial, final int
	var values [3][]int

	if x < 1 || y < 1 {
		panic("A entrada deve ser sempre maior que zero")
	}

	if x <= y {
		initial, final = x, y

	} else {
		final, initial = x, y
	}

	for i := initial; i <= final; i++ {
		for j := i; j <= final; j++ {
			values[0] = append(values[0], j*i)
			values[1] = append(values[1], i)
			values[2] = append(values[2], j)
		}
	}

	return values
}

// PrintResult...
func PrintResult(values [3][]int) {
	minor, major := getMinorMajorPalindromes(values)

	for i := 0; i < len(values[0]); i++ {
		if values[0][i] == minor {
			fmt.Printf("Menor: %d, | Operadores: (%d, %d)\n", minor, values[1][i], values[2][i])
		}
		if values[0][i] == major {
			fmt.Printf("Maior: %d, | Operadores: (%d, %d)\n", major, values[1][i], values[2][i])
		}
	}

}

// isPalindrome...
func isPalindrome(val int) bool {
	input := val
	var remainder int
	result := 0
	for val > 0 {
		remainder = val % 10
		result = (result * 10) + remainder
		val = val / 10
	}
	return input == result
}

// getMinorMajorPalindromes...
func getMinorMajorPalindromes(values [3][]int) (int, int) {
	var minor, major int

	for _, item := range values[0] {
		if isPalindrome(item) {
			if minor == 0 || item < minor {
				minor = item
			}
			if item > major {
				major = item
			}
		}
	}

	return minor, major
}
