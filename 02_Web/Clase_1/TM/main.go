package main

import (
	"encoding/json"
	"os"

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

func main() {
	r := gin.Default()

	r.POST("/hello", helloHandler())
	r.GET("/transactions", getAll())

	r.Run()
}

// helloHandler returns a greeting message with the name of the user
func helloHandler() gin.HandlerFunc {
	type request struct {
		Nombre string `json:"nombre"`
	}
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Hola " + req.Nombre})
	}
}

// getAll returns all the transactions
func getAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactions, err := readFromFile()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, transactions)
	}
}

func readFromFile() ([]Transaction, error) {
	fileData, err := os.ReadFile("transactions.json")
	if err != nil {
		return nil, err
	}
	var transactions []Transaction
	err = json.Unmarshal(fileData, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
