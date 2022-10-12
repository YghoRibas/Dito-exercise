package EX01

import (
	"testing"
)

func TestShouldTestASubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,5}

	var expectedResult = true;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}
}

func TestShouldTestANonSubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,5,3}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro msg")
	}
}

func TestShouldTestValuesOutOfPrincipalArray(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldTestAVoidArrayPrincipalListAsArgument(t *testing.T) {
	var pTest = []int {}
	var sTest = []int {2,4,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldTestAVoidArraySubsequenceListAsArgument(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldTestIfSubsequenceListIsBiggerThanPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,2,3,4,5,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldSumMoreThanOneValuesOutOfPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,7,8,2}

	var expectedResult = 15;

	finalResult := SumNumberOutOfPrincipalList(pTest,sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldSumOnlyOneValuesOutOfPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,7,5,2}

	var expectedResult = 7;

	finalResult := SumNumberOutOfPrincipalList(pTest,sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}

func TestShouldSumNoOneValuesOutOfPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,4,5,2}

	var expectedResult = 0;

	finalResult := SumNumberOutOfPrincipalList(pTest,sTest)

	if finalResult != expectedResult {
		t.Error("erro")
	}	
}