package main

// "github.com/umpc/go-sortedmap"
// "github.com/umpc/go-sortedmap/asc"

import(
	"fmt"
)

func main() {
	fruits := map[string]int {
		"orange": 112,
		"apple": 200,
		"strawberry": 56,
		"banana":174,
	}

	fruits["pineapple"] = 873
	fmt.Println(fruits)

	fruits = removeElementsUnder100(fruits)
	fmt.Println(fruits)

	// fruits := sortedmap.New(5, asc.Int)
	// fruits.Insert("Laranja", 112)
	// fruits.Insert("Abacaxi", 867)
	// fruits.Insert("Manga", 316)
	// fruits.Insert("Melancia", 1343)
	// fruits.Insert("uva", 46)
	// fruits.Insert("Morango",52)

	// fruits.BoundedDelete(0,100)
	// fruits.BoundedDelete(0,100)

 	// fmt.Println("map: ", fruits)
}

func removeElementsUnder100(m map[string]int) map[string]int {
	for key := range m {
		if m[key] < 100 {
			delete(m, key)
		} 
	}
	return m
} 