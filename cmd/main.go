package main

import (
	"d-api/cmd/endpoints/drinks"
	"d-api/cmd/endpoints/foods"
	"d-api/cmd/endpoints/orders"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
)

func main() {
	router := gin.Default()
	router.GET("/foods", foods.GetFoods)
	router.POST("/foods", foods.PostFoods)
	router.GET("/foods/food/:name", foods.GetFoodsByName)
	router.GET("/foods/:id", foods.GetFoodByID)
	router.GET("/drinks", drinks.GetDrinks)
	router.POST("/drinks", drinks.PostDrinks)
	router.GET("/drinks/:id", drinks.GetDrinksByID)
	router.GET("/drinks/drink/:name", drinks.GetDrinksByName)
	router.GET("/orders", orders.GetOrders)
	router.POST("/orders", orders.PostOrders)
	router.GET("/orders/:id", orders.GetOrderByID)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
