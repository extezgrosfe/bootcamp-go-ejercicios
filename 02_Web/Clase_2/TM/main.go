package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

var Transactions []*Transaction

func main() {
	r := gin.Default()

	r.POST("/transactions", create())
	r.GET("/transactions", getAll())

	r.Run()
}

//create creates a new transaction
func create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "secret" {
			c.String(401, "no tiene permisos para realizar la petición solicitada")
			return
		}

		var transaction Transaction
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// validate that every field is not empty
		if transaction.Codigo == "" {
			c.JSON(400, gin.H{"error": createErrorString("codigo")})
			return
		}
		if transaction.Moneda == "" {
			c.JSON(400, gin.H{"error": createErrorString("moneda")})
			return
		}
		if transaction.Monto == 0 {
			c.JSON(400, gin.H{"error": createErrorString("monto")})
			return
		}
		if transaction.Emisor == "" {
			c.JSON(400, gin.H{"error": createErrorString("emisor")})
			return
		}
		if transaction.Receptor == "" {
			c.JSON(400, gin.H{"error": createErrorString("receptor")})
			return
		}
		if transaction.Fecha == "" {
			c.JSON(400, gin.H{"error": createErrorString("fecha")})
			return
		}

		transaction.Id = getLastId() + 1
		Transactions = append(Transactions, &transaction)

		c.JSON(201, transaction)
	}
}

// getAll returns all the transactions
func getAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "secret" {
			c.String(401, "no tiene permisos para realizar la petición solicitada")
			return
		}

		c.JSON(200, Transactions)
	}
}

// getLastId returns the last id
func getLastId() int {
	if len(Transactions) == 0 {
		return 0
	}
	return Transactions[len(Transactions)-1].Id
}

// createError creates an error with the given format
func createErrorString(field string) string {
	errorString := "el campo %s es requerido"
	return fmt.Sprintf(errorString, field)
}
