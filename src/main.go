package main

//"github.com/YghoRibas/Dito-exercise/EX01"
//"github.com/YghoRibas/Dito-exercise/EX02"

import(
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"github.com/YghoRibas/Dito-exercise/EX04"
)

func main() {
// 	var p = []int {1,2,4,6,3}
// 	var s = []int {1,4,3}

// 	showResult := EX01.IsSubsequence(p,s)
// 	fmt.Println("result:", showResult.ItsSubsequence, "    out of array sum:", showResult.SumOutOfPrincipal)

	// milk := EX02.Milk{
	// 	Name: "Milk",
	// 	Weight: 1.025,
	// 	Count: 6,
	// 	OutOfDate: time.Now(),
	// }

	// rice := EX02.Rice{
	// 	Name: "Rice",
	// 	Weight: 1.025,
	// 	Count: 6,
	// 	OutOfDate: time.Now(),
	// }

	// bottle := EX02.Bottle{
	// 	Name: "Bottle",
	// 	Weight: 1.025,
	// 	Count: 6,
	// 	OutOfDate: time.Now(),
	// }

	// market := EX02.Market{
	// 	Name: "testM",
	// 	Address: "Av.test 300",
	// 	Phone: 123456789,
	// 	Products: []EX02.Product{
	// 		milk,
	// 		rice,
	// 		bottle,
	// 	},
	// }
	
	// fmt.Println(market)
	
	// batata := EX03.Fries{
	// 	Name:"Batata frita",
	// 	Ingredients: []string{
	// 		"Batata",
	// 		"Oleo",
	// 	},
	// 	PrepTime: 16,
	// }
	// cooker := EX03.Cooker{}

	// pizza := EX03.Pizza{
	// 	Name:"Pizza de calabresa",
	// 	Ingredients: []string{
	// 		"Calabresa",
	// 		"Queijo",
	// 		"Carvao",
	// 	},
	// 	PrepTime: 43,
	// }

	// cooker.Cook(batata)
	// cooker.Cook(pizza)

	pa := EX04.PersonAddress{
		Street: "av. a",
		Number: 30817,
		Country: "Brazil",
		PostalCode: 123412431,
	}

	p := EX04.Person{
		Name: "Andre",
		Age: 31,
		Address: pa,
		Dependents: []string{},
	}

	body, _ := json.Marshal(p)
    payload := bytes.NewBuffer(body)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", payload)
    if err != nil {
        log.Fatalln(err)
    }

	body, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    log.Printf("%s", body)
}