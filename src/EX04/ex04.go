package EX04

type Person struct {
	Name string `json:"name"`
	Age	int		 `json:"age"`
	Gender	*string `json:"gender"`
	Address PersonAddress `json:"address"`
	Dependents []string  `json:"dependents"`
}

type PersonAddress struct {
	Street string `json:"street"`
	Number int `json:"number"`
	Country string `json:"country"`
	PostalCode int `json:"postal_code"`
}