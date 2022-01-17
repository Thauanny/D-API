package main

import (
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	"net/http"
)

type food struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       [2]string `json:"image"`
	Price       float32   `json:"price"`
}

type drink struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       [2]string `json:"image"`
	Price       float32   `json:"price"`
}

var foods = []food{
	{ID: "1", Name: "Chicken parm", Image: [2]string{
		"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"},
		Description: "simple chicken parm",
		Price:       56.99},
	{ID: "2", Name: "Salad", Image: [2]string{"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"}, Description: "simple salad", Price: 17.99},
	{ID: "3", Name: "Lasagna", Image: [2]string{"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"}, Description: "simple lasagna", Price: 39.99},
}

var drinks = []drink{
	{ID: "1", Name: "Chicken parm", Image: [2]string{
		"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"},
		Description: "simple chicken parm",
		Price:       56.99},
	{ID: "2", Name: "Salad", Image: [2]string{"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"}, Description: "simple salad", Price: 17.99},
	{ID: "3", Name: "Lasagna", Image: [2]string{"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"}, Description: "simple lasagna", Price: 39.99},
}

func getFoods(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, foods)
}

func postFoods(c *gin.Context) {
	var newFood food

	if err := c.BindJSON(&newFood); err != nil {
		return
	}

	foods = append(foods, newFood)
	c.IndentedJSON(http.StatusCreated, newFood)
}

func getFoodByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range foods {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
}

func getDrinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drinks)
}

func postDrinks(c *gin.Context) {
	var newDrink drink

	if err := c.BindJSON(&newDrink); err != nil {
		return
	}

	drinks = append(drinks, newDrink)
	c.IndentedJSON(http.StatusCreated, newDrink)
}

func getDrinksByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range drinks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "drink not found"})
}

func main() {
	router := gin.Default()
	router.GET("/foods", getFoods)
	router.POST("/foods", postFoods)
	router.GET("/foods/:id", getFoodByID)
	router.GET("/drinks", getDrinks)
	router.POST("/drinks", postDrinks)
	router.GET("/drinks/:id", getDrinksByID)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
