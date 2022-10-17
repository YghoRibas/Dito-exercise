package EX01

// SubsequenceResult ...
type SubsequenceResult struct {
	ItsSubsequence 		bool
	SumOutOfPrincipal	int
}

// IsSubsequence ...
func IsSubsequence(listP []int, listS []int) SubsequenceResult {
	var sResult SubsequenceResult
	listPSize := len(listP)
	listSSize := len(listS)
	var indexCounter int = 0
	var auxiliar = make([]int, listSSize)

	sResult.ItsSubsequence = true

	if !verifyArraySizeIsValid(listPSize,listSSize) {
		sResult.ItsSubsequence = false
		sResult.SumOutOfPrincipal = sumNumberOutOfPrincipalList(listP,listS)
		return sResult
	}

	for i:= 0; i < listSSize; i++ {
		for j:= indexCounter; j < listPSize; j++ {
			if listP[j] == listS[i] {
				auxiliar[i] = listS[i]
				indexCounter = j + 1
				break
			}
			if j == listPSize - 1 {
				sResult.ItsSubsequence = false
			}
		}
		if auxiliar[i] != listS[i] {
			sResult.ItsSubsequence = false
		}
		if !sResult.ItsSubsequence {
			break
		}
	}

	sResult.SumOutOfPrincipal = sumNumberOutOfPrincipalList(listP,listS)
	
	return sResult
}

func verifyArraySizeIsValid(listPSize int, listSSize int) bool {
	return listSSize <= listPSize && listPSize > 0 && listSSize > 0
}

func sumNumberOutOfPrincipalList(listP []int, listS []int) int {
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
		if listPSize == 0 {
			outOfArrayCounter += listS[i]
		}
	}

	return outOfArrayCounter
}
