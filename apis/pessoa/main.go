package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Pessoa struct {
	Id    int64  `json:id`
	Nome  string `json:nome`
	CPF   string `json:cpf`
	Idade int    `json:idade`
	Email string `json:email`
}

type Nome struct {
	Nome string `json:nome`
}

type CPF struct {
	CPF string `json:cpf`
}

type Erro struct {
	Mensagem string `json:mensagem`
}

type IdadeEmail struct {
	Idade int    `json:idade`
	Email string `json:email`
}

func main() {
	router := gin.Default()
	router.POST("/pessoa", postPessoa)
	router.GET("/health", getHealth)
	router.GET("/ready", getReady)

	router.Run(":8080")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func postPessoa(c *gin.Context) {

	pessoa := Pessoa{}

	var nome Nome
	var erro Erro
	var cpf CPF
	var idadeEmail IdadeEmail

	MYSQL_SERVER := os.Getenv("MYSQL_SERVER")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	//Obter um nome
	responseNome, err := http.Get("http://api-nome/new")

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseNome.Body == nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	responseNomeData, err := ioutil.ReadAll(responseNome.Body)

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseNome.StatusCode == 200 {
		json.Unmarshal(responseNomeData, &nome)
	} else {
		json.Unmarshal(responseNomeData, &erro)
		c.IndentedJSON(responseNome.StatusCode, erro)
		return
	}

	pessoa.Nome = nome.Nome

	//Obter um CPF
	responseCPF, err := http.Get("http://api-cpf/new")

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseNome.Body == nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	responseCPFData, err := ioutil.ReadAll(responseCPF.Body)

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseCPF.StatusCode == 200 {
		json.Unmarshal(responseCPFData, &cpf)
	} else {
		json.Unmarshal(responseCPFData, &erro)
		c.IndentedJSON(responseCPF.StatusCode, erro)
		return
	}

	pessoa.CPF = cpf.CPF

	//Obter uma Idade e E-mail
	responseIdadeEmail, err := http.Get("http://api-idade-email/new")

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseNome.Body == nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	responseIdadeEmailData, err := ioutil.ReadAll(responseIdadeEmail.Body)

	if err != nil {
		erro = Erro{Mensagem: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, erro)
		return
	}

	if responseIdadeEmail.StatusCode == 200 {
		json.Unmarshal(responseIdadeEmailData, &idadeEmail)
	} else {
		json.Unmarshal(responseIdadeEmailData, &erro)
		c.IndentedJSON(responseIdadeEmail.StatusCode, erro)
		return
	}

	pessoa.Idade = idadeEmail.Idade
	pessoa.Email = idadeEmail.Email

	db, err := sql.Open("mysql", MYSQL_USER+":"+MYSQL_PASSWORD+"@tcp("+MYSQL_SERVER+":"+MYSQL_PORT+")/"+MYSQL_DATABASE)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO pessoa(nome, cpf, idade, email) VALUES ('" + pessoa.Nome + "', '" + pessoa.CPF + "', " + strconv.Itoa(pessoa.Idade) + ", '" + pessoa.Email + "')"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	pessoa.Id = lastId

	c.IndentedJSON(http.StatusOK, pessoa)

}

func getReady(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
