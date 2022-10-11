package EX01

import (
	"testing"
)

func TestShouldTestASubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,5}

	var expectedTest = true;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro")
	}
}

func TestShouldTestANonSubsequence(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,5,3}

	var expectedTest = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro msg")
	}
}

func TestShouldTestValuesOutOfPrincipalArray(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {2,4,6}

	var expectedTest = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro")
	}	
}

func TestShouldTestAVoidArrayPrincipalListAsArgument(t *testing.T) {
	var pTest = []int nil
	var sTest = []int {2,4,6}

	var expectedTest = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro")
	}	
}

func TestShouldTestAVoidArraySubsequenceListAsArgument(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int nil

	var expectedTest = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro")
	}	
}

func TestShouldTestIfSubsequenceListIsBiggerThanPrincipalList(t *testing.T) {
	var pTest = []int {1,2,3,4,5}
	var sTest = []int {1,2,3,4,5,6}

	var expectedTest = false;

	finalResult := IsSubsequence(pTest, sTest)

	if finalResult != expectedTest {
		t.Error("erro")
	}	
}