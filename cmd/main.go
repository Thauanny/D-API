package main

import (
	"d-api/cmd/endpoints/drinks"
	"d-api/cmd/endpoints/foods"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
)

func main() {
	router := gin.Default()
	router.GET("/foods", foods.GetFoods)
	router.POST("/foods", foods.PostFoods)
	router.GET("/foods/:id", foods.GetFoodByID)
	router.GET("/drinks", drinks.GetDrinks)
	router.POST("/drinks", drinks.PostDrinks)
	router.GET("/drinks/:id", drinks.GetDrinksByID)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
