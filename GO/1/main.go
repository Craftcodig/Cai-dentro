package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Estrutura do que uma pessoa tem de dados
type pessoa struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Lista de pessoas fixas guardadas em memoria
var listaPessoasFixas = []pessoa{
	{ID: "1", Nome: "Jose", Email: "josedascouces@gmail.com"},
	{ID: "2", Nome: "Maria", Email: "mariadashortas@gmail.com"},
	{ID: "3", Nome: "Joao", Email: "joao@bbb"},
	{ID: "4", Nome: "Milton", Email: "miltola@empresadogas"},
}

// lista de pessoas dinamicas em memoria
var listaPessoasDinamicas = []pessoa{}

// Programa principal do Go
func main() {
	rotas := gin.Default()
	rotas.GET("/listaPessoas", getListaPessoas)
	rotas.POST("/salvaPessoa", salvaPessoa)

	rotas.Run("localhost:8080")
}

// Serve para listar pessoas baseado no numeor da lista
func getListaPessoas(c *gin.Context) {
	numeroLista := c.Query("id")

	var lista []pessoa

	// TODO em vez de if pode ser um switch
	if numeroLista == "1" {
		lista = listaPessoasFixas
	} else {
		if numeroLista == "2" {
			lista = listaPessoasDinamicas
		} else {
			c.IndentedJSON(http.StatusBadRequest, "Essa lista nao existe!")
			return
		}
	}
	c.IndentedJSON(http.StatusOK, lista)
}

// Salvar pessoas na lista dinamica
func salvaPessoa(c *gin.Context) {
	pessoa := pessoa{}

	if err := c.BindJSON(&pessoa); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO validar se os dados estão preenchidos (restrições)

	// Salva na lista
	listaPessoasDinamicas = append(listaPessoasDinamicas, pessoa)

	c.IndentedJSON(http.StatusOK, "Salvou com sucesso!")
}