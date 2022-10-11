package main

import(
	"fmt"
	"context"

	"github.com/YghoRibas/Dito-exercise/EX01"
)

func main() {
	var p = []int {1,2,3,4,5}
	var s = []int {2,5,4}
	fmt.Println(IsSubsequence(p, s))
}