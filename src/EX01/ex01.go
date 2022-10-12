package EX01

// import (
// 	"fmt"
// )

func IsSubsequence(listP []int, listS []int) bool {
	listPSize := len(listP)
	listSSize := len(listS)
	var auxiliar = make([]int, listSSize)
	var indexCounter int = 0

	if listSSize > listPSize || (listPSize < 1 || listSSize < 1) {
		return false
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
			return false
		}
	}
	
	return true
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
