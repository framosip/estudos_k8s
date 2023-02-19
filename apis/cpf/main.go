package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CPF struct {
	CPF string `json:cpf`
}

func main() {
	router := gin.Default()
	router.GET("/new", getNew)
	router.GET("/health", getHealth)

	router.Run(":8080")
}

func getNew(c *gin.Context) {

	cpf := ""

	for i := 1; i <= 14; i++ {
		if i == 4 || i == 8 {
			cpf = cpf + "."
		} else if i == 11 {
			cpf = cpf + "-"
		} else {
			cpf = cpf + strconv.Itoa(rand.Intn(9-0)+0)
		}
	}

	c.IndentedJSON(http.StatusOK, CPF{CPF: cpf})
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
