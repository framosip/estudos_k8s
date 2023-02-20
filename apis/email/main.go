package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Email string `json:email`
}

var provedores = []string{"@gmail.com", "@hotmail.com", "@yahoo.com"}

func main() {
	router := gin.Default()
	router.GET("/new", getNew)
	router.GET("/health", getHealth)
	router.GET("/ready", getReady)

	router.Run(":8080")
}

func getNew(c *gin.Context) {
	idProvedor := rand.Intn(2-0) + 0
	idEmail := rand.Intn(20-0) + 0
	c.IndentedJSON(http.StatusOK, Email{Email: "email_" + strconv.Itoa(idEmail) + provedores[idProvedor]})
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func getReady(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
