package drinks

import (
	"d-api/cmd/entities"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var drinks = []entities.Drink{
	{ID: "1", Name: "Coca cola", Image: [2]string{
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSw1d6ThcGR9k6TRjcsOA8hdwSRYoztNDR0HtORRH5s1Q&s",
		"https://www.designi.com.br/images/preview/10000392.jpg"},
		Description: "coca cola",
		Time:        "00:02:00",
		Price:       6.99},
	{ID: "2", Name: "Pepsi", Image: [2]string{
		"https://a-static.mlcdn.com.br/470x352/pepsi-lata-350ml-pack-12-unidades/distribuidoravitalli/715p/e0703c7cc57cebc07283f814319afc46.jpg",
		"https://apoioentrega.vteximg.com.br/arquivos/ids/459556-1000-1000/16a971a96ac5fc6b02ac8f5568e0ad8a_refrigerante-pepsi-lata-350-ml---refrig-pepsi-350ml-lt-cola---1-un_lett_1.jpg?v=637305880434630000"},
		Description: "pepsi",
		Time:        "00:02:00",
		Price:       6.99},
	{ID: "3", Name: "Guarana", Image: [2]string{
		"https://static.paodeacucar.com/img/uploads/1/477/19514477.jpg",
		"https://m.media-amazon.com/images/I/811VnQ45oHS._AC_SY606_.jpg"},
		Description: "guarana",
		Time:        "00:02:00",
		Price:       6.99},
	{ID: "4", Name: "Fanta", Image: [2]string{
		"https://casafiesta.fbitsstatic.net/img/p/refrigerante-fanta-laranja-lata-350ml-67109/233977.jpg?w=420&h=420&v=no-change&qs=ignore",
		"https://www.imigrantesbebidas.com.br/bebida/images/products/full/1990-refrigerante-fanta-laranja-lata-220ml.jpg"},
		Description: "fanta",
		Time:        "00:02:00",
		Price:       6.99},
	{ID: "5", Name: "Fanta uva", Image: [2]string{
		"https://trimais.vteximg.com.br/arquivos/ids/1028659-1000-1000/foto_original.jpg?v=637716210071430000",
		"https://deskontao.agilecdn.com.br/28059_1.jpg"},
		Description: "fanta uva",
		Time:        "00:02:00",
		Price:       6.99},
	{ID: "6", Name: "Sprite", Image: [2]string{
		"https://www.drogariaminasbrasil.com.br/media/product/61d/refrigerante-sprite-lata-350ml-123.jpg",
		"https://static.paodeacucar.com/img/uploads/1/904/11840904.jpg"},
		Description: "sprite",
		Time:        "00:02:00",
		Price:       6.99},
}

func GetDrinks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drinks)
}

func PostDrinks(c *gin.Context) {
	var newDrink entities.Drink

	if err := c.BindJSON(&newDrink); err != nil {
		return
	}

	drinks = append(drinks, newDrink)
	c.IndentedJSON(http.StatusCreated, newDrink)
}

func GetDrinksByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range drinks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "drink not found"})
}

func GetDrinksByName(c *gin.Context) {
	name := c.Param("name")

	for _, a := range drinks {
		if strings.Contains(a.Name, name) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
}
