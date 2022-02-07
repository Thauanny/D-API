package orders

import (
	"d-api/cmd/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

var orders = []entities.Order{
	{Number: "1", Type: "Chicken parm", Items: []entities.Item{}}}

func GetOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)
}

func PostOrders(c *gin.Context) {
	var newOrder entities.Order

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, newOrder)
}

func GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range orders {
		if a.Number == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
}
