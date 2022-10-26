package main

//"github.com/YghoRibas/Dito-exercise/EX01"
//"github.com/YghoRibas/Dito-exercise/EX02"

import(
	"net/http"
	"github.com/YghoRibas/Dito-exercise/EX04"
	"github.com/gin-gonic/gin"
)

var pa = EX04.PersonAddress{
	Street: "av. a",
	Number: 30817,
	Country: "Brazil",
	PostalCode: 123412431,
}

var p = EX04.Person{
	Name: "Andre",
	Age: 31,
	Address: pa,
	Dependents: []string{},
}

func main() {
	router := gin.Default()
    router.POST("/person", postPerson)
	router.GET("/person", getPerson)

    router.Run("localhost:3333")
}

func postPerson(c *gin.Context) {
    var newDep string

    if err := c.BindJSON(&newDep); err != nil {
        return
    }

    p.Dependents = append(p.Dependents, newDep)
    c.IndentedJSON(http.StatusCreated, newDep)
}

func getPerson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, p)
}	
