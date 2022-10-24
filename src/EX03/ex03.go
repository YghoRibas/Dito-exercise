package EX03

import(
	"fmt"
)

type Cooker struct {

}

func (c Cooker) Cook(d Dish) {
	d.CookDish()
}

type Dish interface {
	CookDish()
}	

type Pizza struct {
	Name string
	Ingredients []string
	PrepTime int //minutes
}

func (p Pizza) CookDish() {
	fmt.Printf("COZINHAR PRATO %s PELOS PRÓXIMOS %d MINUTOS \n", p.Name, p.PrepTime)
}

type Fries struct {
	Name string
	Ingredients []string
	PrepTime int //minutes
}

func (f Fries) CookDish() {
	fmt.Printf("COZINHAR PRATO %s PELOS PRÓXIMOS %d MINUTOS \n", f.Name, f.PrepTime)
}