package main

import (
	"encoding/json"
	"os"
	"strconv"

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

type TransactionFilter struct {
	Codigo   string  `form:"codigo"`
	Moneda   string  `form:"moneda"`
	Monto    float64 `form:"monto"`
	Emisor   string  `form:"emisor"`
	Receptor string  `form:"receptor"`
	Fecha    string  `form:"fecha"`
}

func main() {
	r := gin.Default()

	r.POST("/hello", helloHandler())
	r.GET("/transactions", getAll())
	r.GET("/transactions/:id", getById())

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

		// apply filters
		var filter TransactionFilter
		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		transactions = transactionFilter(transactions, filter)

		c.JSON(200, transactions)
	}
}

// getById returns the transaction with the given id
func getById() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		// id to int
		id, err := strconv.Atoi(paramId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		transactions, err := readFromFile()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		for _, transaction := range transactions {
			if transaction.Id == id {
				c.JSON(200, transaction)
				return
			}
		}

		c.JSON(404, gin.H{"error": "Transaction not found"})
	}
}

// transactionFilter returns the transactions filtered by the given filter
func transactionFilter(transactions []Transaction, filter TransactionFilter) []Transaction {
	filteredTransactions := []Transaction{}
	for _, transaction := range transactions {
		if filter.Codigo != "" && filter.Codigo != transaction.Codigo {
			continue
		}
		if filter.Moneda != "" && filter.Moneda != transaction.Moneda {
			continue
		}
		if filter.Monto != 0 && filter.Monto != transaction.Monto {
			continue
		}
		if filter.Emisor != "" && filter.Emisor != transaction.Emisor {
			continue
		}
		if filter.Receptor != "" && filter.Receptor != transaction.Receptor {
			continue
		}
		if filter.Fecha != "" && filter.Fecha != transaction.Fecha {
			continue
		}
		filteredTransactions = append(filteredTransactions, transaction)
	}

	return filteredTransactions
}

// readFromFile reads the transactions from the file
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
