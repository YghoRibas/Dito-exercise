package main

// "github.com/umpc/go-sortedmap"
// "github.com/umpc/go-sortedmap/asc"

import (
	"fmt"

	"github.com/YghoRibas/Dito-exercise/src/EX05"
)

func main() {

	minor, major := EX05.GetMinorMajorPalindromesBetweenRange(1, 4)

	fmt.Printf("Menor: %d, Operadores: ", minor.Product)
	for i := 0; i < len(minor.PairX); i++ {
		fmt.Printf("(%d , %d)", minor.PairX[i], minor.PairY[i])
	}

	fmt.Printf("Maior: %d, Operadores: ", major.Product)
	for index := range major.PairY {
		fmt.Printf("(%d , %d)", major.PairX[index], major.PairY[index])
	}

}
