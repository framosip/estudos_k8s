package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Nome struct {
	Nome string `json:nome`
}

var nomes = []string{"Miguel", "Sophia", "Davi", "Alice", "Arthur", "Julia", "Pedro", "Isabella", "Gabriel", "Manuela"}
var sobrenomes = []string{"Andrade", "Barbosa", "Barros", "Borges", "Freitas", "Gonçalves", "Lima", "Monteiro", "Oliveira", "Silva"}

func main() {
	router := gin.Default()
	router.GET("/new", getNew)
	router.GET("/health", getHealth)
	router.GET("/ready", getReady)

	router.Run(":8080")
}

func getNew(c *gin.Context) {
	idNome := rand.Intn(9-0) + 0
	idSobrenome := rand.Intn(9-0) + 0
	c.IndentedJSON(http.StatusOK, Nome{Nome: nomes[idNome] + " " + sobrenomes[idSobrenome]})
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func getReady(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
