package main

import(
	"fmt"

	"github.com/YghoRibas/Dito-exercise/EX01"
)

func main() {
	var p = []int {1,2,4,6}
	var s = []int {1,4,5}

	showResult := EX01.IsSubsequence(p,s)
	fmt.Println("result:", showResult.ItsSubsequence, "    out of array sum:", showResult.SumOutOfPrincipal)
}