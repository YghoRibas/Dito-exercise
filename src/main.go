package main

//"github.com/YghoRibas/Dito-exercise/EX01"

import(
	"fmt"
	"time"
	"github.com/YghoRibas/Dito-exercise/EX02"
)

func main() {
// 	var p = []int {1,2,4,6,3}
// 	var s = []int {1,4,3}

// 	showResult := EX01.IsSubsequence(p,s)
// 	fmt.Println("result:", showResult.ItsSubsequence, "    out of array sum:", showResult.SumOutOfPrincipal)

	milk := EX02.Milk{
		Name: "Milk",
		Weight: 1.025,
		Count: 6,
		OutOfDate: time.Now(),
	}

	rice := EX02.Rice{
		Name: "Rice",
		Weight: 1.025,
		Count: 6,
		OutOfDate: time.Now(),
	}

	bottle := EX02.Bottle{
		Name: "Bottle",
		Weight: 1.025,
		Count: 6,
		OutOfDate: time.Now(),
	}

	market := EX02.Market{
		Name: "testM",
		Address: "Av.test 300",
		Phone: 123456789,
		Products: []EX02.Product{
			milk,
			rice,
			bottle,
		},
	}
	
	fmt.Println(market)
}