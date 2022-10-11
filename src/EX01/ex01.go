package EX01

func IsSubsequence(listP []int, listS []int) bool {
	listPSize := len(listP)
	listSSize := len(listS)
	var aux []int

	if listSSize > listPSize {
		return false
	}

	for i := 0; i < listSSize; i++ {
		for j := 0 + i; j < listPSize; j++ {
			if listP[j] == listS[i] {
				aux[i] = listS[i]
			}
		}
	} 

	return aux[] == listS
}