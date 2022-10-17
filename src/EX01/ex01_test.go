package EX01

import (
	"testing"
)

func TestShouldBeASubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,5}

	var expectedResult = true;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro")
	}
	if finalResult.SumOutOfPrincipal != 0 {
		t.Error("erro")
	}
}

func TestShouldBeANonSubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,5,3}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro msg")
	}
	if finalResult.SumOutOfPrincipal != 0 {
		t.Error("erro")
	}
}

func TestShouldTestValuesOutOfPrincipalArray(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro")
	}
	if finalResult.SumOutOfPrincipal != 6 {
		t.Error("erro")
	}	
}

func TestShouldTestAVoidArrayPrincipalListAsArgument(t *testing.T) {
	var pTest = []int {}
	var sTest = []int {2,4,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro")
	}
	if finalResult.SumOutOfPrincipal != 12 {
		t.Error("erro")
	}		
}

func TestShouldTestAVoidArraySubsequenceListAsArgument(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro")
	}
	if finalResult.SumOutOfPrincipal != 0 {
		t.Error("erro")
	}		
}

func TestShouldTestIfSubsequenceListIsBiggerThanPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,2,3,4,5,6}

	var expectedResult = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult.ItsSubsequence != expectedResult {
		t.Error("erro")
	}
	if finalResult.SumOutOfPrincipal != 6 {
		t.Error("erro")
	}		
}