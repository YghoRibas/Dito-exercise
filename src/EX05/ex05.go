package EX05

import "fmt"

// MultValuesBetween...
func MultValuesBetween(x, y int) [3][]int {
	var initial, final int
	var values [3][]int

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
	resultOperators := getPalindromesAndOperators(values)

	fmt.Println("Menor: ", minor, " | Maior: ", major)
	for i := 0; i < len(resultOperators[0]); i++ {
		fmt.Printf("Resultado: %d | Operadores: (%d , %d)\n", resultOperators[0][i], resultOperators[1][i], resultOperators[2][i])
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

func getPalindromesAndOperators(values [3][]int) [3][]int {
	var multOperators [3][]int

	for i := 0; i < len(values[0]); i++ {
		if isPalindrome(values[0][i]) {
			multOperators[0] = append(multOperators[0], values[0][i])
			multOperators[1] = append(multOperators[1], values[1][i])
			multOperators[2] = append(multOperators[2], values[2][i])
		}
	}

	return multOperators
}
