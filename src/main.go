package main

//"github.com/YghoRibas/Dito-exercise/EX01"
//"github.com/YghoRibas/Dito-exercise/EX02"

import(
	"net/http"
	"github.com/YghoRibas/Dito-exercise/EX04"
	"github.com/gin-gonic/gin"
)

var p = EX04.Person{}

func main() {
	router := gin.Default()
    router.POST("/person", postPerson)
	router.GET("/person", getPerson)

    router.Run("localhost:3000")
}

func postPerson(c *gin.Context) {
    var newPerson EX04.Person

    if err := c.BindJSON(&newPerson); err != nil {
        return
    }

    p.Name = newPerson.Name
	p.Age = newPerson.Age
	p.Address = newPerson.Address
	p.Dependents = newPerson.Dependents
    c.IndentedJSON(http.StatusCreated, newPerson)
}

func getPerson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, p)
}
