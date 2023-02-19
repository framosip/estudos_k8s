package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IdadeEmail struct {
	Idade int    `json:idade`
	Email string `json:email`
}

type Email struct {
	Email string `json:email`
}

type Erro struct {
	Mensagem string `json:mensagem`
}

func main() {
	router := gin.Default()
	router.GET("/new", getNew)
	router.GET("/health", getHealth)

	router.Run(":8080")
}

func getNew(c *gin.Context) {

	var erro Erro
	var email Email

	//Obter um E-mail
	responseEmail, err := http.Get("http://api-email/new")

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseEmail.Body == nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	responseEmailData, err := ioutil.ReadAll(responseEmail.Body)

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseEmail.StatusCode == 200 {
		json.Unmarshal(responseEmailData, &email)
	} else {
		json.Unmarshal(responseEmailData, &erro)
		c.IndentedJSON(responseEmail.StatusCode, erro)
		return
	}

	idade := rand.Intn(80-0) + 0
	c.IndentedJSON(http.StatusOK, IdadeEmail{Idade: idade, Email: email.Email})
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
