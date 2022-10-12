package main

import(
	"fmt"

	"github.com/YghoRibas/Dito-exercise/EX01"
)

func main() {
	var p = []int {1,2,4,6}
	var s = []int {1,4,3,7}
	fmt.Println("result:", EX01.IsSubsequence(p, s), "    out of array sum:", EX01.SumNumberOutOfPrincipalList(p, s))
}