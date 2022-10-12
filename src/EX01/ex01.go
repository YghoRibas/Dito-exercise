package EX01

// import (
// 	"fmt"
// )

type SubsequenceResult struct {
	ItsSubsequence 		bool
	SumOutOfPrincipal	int
}

func IsSubsequence(listP []int, listS []int) SubsequenceResult {
	var sResult SubsequenceResult
	listPSize := len(listP)
	listSSize := len(listS)
	var auxiliar = make([]int, listSSize)
	var indexCounter int = 0

	sResult.ItsSubsequence = true

	if listSSize > listPSize || (listPSize < 1 || listSSize < 1) {
		sResult.ItsSubsequence = false
	} 

	for i:= 0; i < listSSize; i++ {
		for j:= indexCounter; j < listPSize; j++ {
			if listP[j] == listS[i] {
				auxiliar[i] = listP[j]
				indexCounter = j + 1
				break
			}
		}
		if auxiliar[i] != listS[i] {
			sResult.ItsSubsequence = false
		}
	}

	sResult.SumOutOfPrincipal = SumNumberOutOfPrincipalList(listP,listS)
	
	return sResult
}

func SumNumberOutOfPrincipalList(listP []int, listS []int) int {
	listPSize := len(listP)
	listSSize := len(listS)
	var outOfArrayCounter int = 0

	for i:= 0; i < listSSize; i++ {
		for j:= 0; j < listPSize; j++ {
			if listP[j] == listS[i] {
				break
			} 
			if j == listPSize - 1 {
				outOfArrayCounter += listS[i]
			}
		}
	}
	return outOfArrayCounter
}
