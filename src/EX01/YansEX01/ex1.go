package src

// IsSubsequenceOut represents the output of IsSubsequence.
type IsSubsequenceOut struct {
	IsSubsequence bool
	SumOfValues   int
}

// IsSubsequence checks if secondary list is a subsequence of principal list.
func IsSubsequence(principalList []int, SecondaryList []int) IsSubsequenceOut {
	var currentIndex int = -1
	var sumOfValues int
	var isSub bool = true

	if len(SecondaryList) == 0 {
		return IsSubsequenceOut{}
	}

	for _, itemSecondary := range SecondaryList {
		isInList, index := checkIfElementIsInList(itemSecondary, principalList)
		if isInList {
			if index > currentIndex {
				currentIndex = index
				continue
			}
		} else {
			sumOfValues += itemSecondary
		}
		isSub = false
	}

	return IsSubsequenceOut{
		IsSubsequence: isSub,
		SumOfValues:   sumOfValues,
	}
}

func checkIfElementIsInList(element int, list []int) (bool, int) {
	for indexList, item := range list {
		if item == element {
			return true, indexList
		}
	}
	return false, -1
}
