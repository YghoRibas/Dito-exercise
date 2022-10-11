package EX01

// import (
// 	"fmt"
// )

func IsSubsequence(listP []int, listS []int) bool {
	listPSize := len(listP)
	listSSize := len(listS)
	var aux []int
	var counter int = 0

	if listSSize > listPSize || (listP == nil || listS == nil) {
		return false
	} 

	for i:= 0; i < listSSize; i++ {
		for j:= counter; j < listPSize; j++ {
			if listP[j] == listS[i] {
				aux[i] = listS[i]
				counter = j + 1
				break
			}
		}
	}
	
	for k := 0; k < listSSize; k++ {
		if aux[k] != listS[k] {
			return false
		}
	}
	
	return true
}