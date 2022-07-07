package main

import (
	"encoding/json"
	"strconv"

	"github.com/Xartyago/XartyaGo-/transaction"
	"github.com/gin-gonic/gin"
)

var Transactions []transaction.Transaction

//This endpoint fill our var Transaction and show them
func getAndUpdateAllTransactions(c *gin.Context) {
	// Query statements
	var realData map[string]string
	possibleQueries := map[string]string{
		"id":       c.Query("Id"),
		"name":     c.Query("Name"),
		"code":     c.Query("Code"),
		"currency": c.Query("Currency"),
		"emisor":   c.Query("Emisor"),
		"receiver": c.Query("Receiver"),
		"amount":   c.Query("Amount"),
		"date":     c.Query("Date"),
	}
	for key, query := range possibleQueries {
		if query != "" {
			realData[key] = query
		}
	}

	//Get without queries
	if len(Transactions) == 0 {
		var transactions []transaction.Transaction
		transactionsJson := transaction.Json()
		json.Unmarshal([]byte(transactionsJson), &transactions)
		*&Transactions = transactions
		c.JSON(200, gin.H{"Transactions": transactions})
	}
}

// Here we go to find a transaction by id
func getSpecificTransaction(c *gin.Context) {
	idToFound, _ := strconv.Atoi(c.Param("id"))
	trcionFounded := Transactions[idToFound-1]
	if trcionFounded.Id > 0 {
		c.JSON(200, trcionFounded)
	} else {
		c.JSON(404, gin.H{"msg": "The tranac not founded"})
	}
}

func main() {
	router := gin.Default()
	transactions := router.Group("/transactions")
	{
		transactions.GET("/", getAndUpdateAllTransactions)
		transactions.GET("/:id", getSpecificTransaction)
	}
	router.Run(":3001")
}
