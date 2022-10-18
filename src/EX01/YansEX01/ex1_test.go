package src

import "testing"

func TestEmptySecondaryList(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != false {
		t.Error("test 1 fails")
	}

	if result.SumOfValues != 0 {
		t.Error("test 1 fails")
	}
}

func TestEmptyPrincipalList(t *testing.T) {
	principalList := []int{}
	secondaryList := []int{1, 2}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != false {
		t.Error("test 2 fails")
	}

	if result.SumOfValues != 3 {
		t.Error("test 2 fails")
	}
}

func TestSubsequenceSuccessCase(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{1, 4}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != true {
		t.Error("test 3 fails")
	}

	if result.SumOfValues != 0 {
		t.Error("test 3 fails")
	}
}

func TestSequencialElementDuplicate(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{1, 1}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != false {
		t.Error("test 4 fails")
	}

	if result.SumOfValues != 0 {
		t.Error("test 4 fails")
	}
}

func TestNonSequencialElementDuplicate(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{1, 2, 1}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != false {
		t.Error("test 5 fails")
	}

	if result.SumOfValues != 2 {
		t.Error("test 5 fails")
	}
}

func TestInvertedSecondaryElements(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{4, 1}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != false {
		t.Error("test 6 fails")
	}

	if result.SumOfValues != 0 {
		t.Error("test 6 fails")
	}
}

func TestUniqueElementInSecondaryList(t *testing.T) {
	principalList := []int{1, 4, 8, -1}
	secondaryList := []int{1}

	result := IsSubsequence(principalList, secondaryList)

	if result.IsSubsequence != true {
		t.Error("test 7 fails")
	}

	if result.SumOfValues != 0 {
		t.Error("test 7 fails")
	}
}
