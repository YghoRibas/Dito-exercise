package EX05

import (
	"testing"
)

func TestShouldReturnMinorMajorPalindromesInRange(t *testing.T) {
	minor, major := GetMinorMajorPalindromesBetweenRange(1, 4)

	expectedMinor := Palindrome{
		Product: 1,
		PairX:   []int{1},
		PairY:   []int{1},
	}

	expectedMajor := Palindrome{
		Product: 9,
		PairX:   []int{3},
		PairY:   []int{3},
	}

	// Compare Minor attributes
	if minor.Product != expectedMinor.Product {
		t.Error("erro")
	}

	for index, item := range minor.PairX {
		if item != expectedMinor.PairX[index] {
			t.Error("erro")
		}
	}

	for index, item := range minor.PairY {
		if item != expectedMinor.PairY[index] {
			t.Error("erro")
		}
	}

	// Compare Major attributes
	if major.Product != expectedMajor.Product {
		t.Error("erro")
	}

	for index, item := range major.PairX {
		if item != expectedMajor.PairX[index] {
			t.Error("erro")
		}
	}

	for index, item := range major.PairY {
		if item != expectedMajor.PairY[index] {
			t.Error("erro")
		}
	}
}
