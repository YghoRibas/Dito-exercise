package EX05

type Palindrome struct {
	Product int
	PairX   []int
	PairY   []int
}

// MultValuesBetween...
func multValuesBetween(x, y int) [3][]int {
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
func GetMinorMajorPalindromesBetweenRange(x, y int) (Palindrome, Palindrome) {

	matrix := multValuesBetween(x, y)
	minor, major := getMinorMajorPalindromes(matrix)

	return minor, major
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
func getMinorMajorPalindromes(values [3][]int) (Palindrome, Palindrome) {
	var minorAux, majorAux int
	var minor, major Palindrome

	for _, item := range values[0] {
		if isPalindrome(item) {
			if minorAux == 0 || item < minorAux {
				minorAux = item

			}
			if item > majorAux {
				majorAux = item
			}
		}
	}

	for index, item := range values[0] {
		if minorAux == item {
			minor.Product = minorAux
			minor.PairX = append(minor.PairX, values[1][index])
			minor.PairX = append(minor.PairY, values[2][index])
		}
		if majorAux == item {
			major.Product = majorAux
			major.PairX = append(major.PairX, values[1][index])
			major.PairY = append(major.PairY, values[2][index])
		}
	}

	return minor, major
}
