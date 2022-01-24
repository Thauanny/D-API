package foods

import (
	"d-api/cmd/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var foods = []entities.Food{
	{ID: "1", Name: "Chicken parm", Image: [2]string{
		"https://www.skinnytaste.com/wp-content/uploads/2012/09/skinny-chicken-parmigiana-550x733.jpg",
		"https://assets.bonappetit.com/photos/5ea9a053093f970009773e21/master/w_1600,c_limit/Chicken-Parm-Pan-Inline.jpg"},
		Description: "simple chicken parm",
		Time:        "00:20:00",
		Price:       56.99},
	{ID: "2", Name: "Salad", Image: [2]string{
		"https://cdn.loveandlemons.com/wp-content/uploads/2019/07/salad.jpg",
		"https://www.onceuponachef.com/images/2019/07/Big-Italian-Salad.jpg"},
		Description: "simple salad",
		Time:        "00:10:00",
		Price:       17.99},
	{ID: "3", Name: "Lasagna", Image: [2]string{
		"https://www.thespruceeats.com/thmb/-YUYSXc4T2H4P8o2JBApzJ3F5rw=/2069x2069/smart/filters:no_upscale()/white-and-red-sauce-lasagna-recipe-995925-hero-1-fb2c2142b39042069f0c50310047256d.jpg",
		"https://www.jessicagavin.com/wp-content/uploads/2017/07/meat-lasagna-1200.jpg"},
		Description: "simple lasagna",
		Time:        "00:30:00",
		Price:       39.99},
	{ID: "4", Name: "Beef", Image: [2]string{
		"https://food.fnr.sndimg.com/content/dam/images/food/fullset/2013/6/19/0/WU0501H_peppercorn-roasted-beef-tenderloin-recipe_s4x3.jpg.rend.hgtvcom.616.462.suffix/1552487731049.jpeg",
		"https://recipetineats.com/wp-content/uploads/2019/05/Marinated-Roast-Beef_7-1.jpg"},
		Description: "beeeef",
		Time:        "00:15:00",
		Price:       49.99},
	{ID: "5", Name: "Roasted Chicken", Image: [2]string{
		"https://cafedelites.com/wp-content/uploads/2017/12/Perfect-Juicy-Roast-Chicken-IMAGE-28.jpg",
		"https://tmbidigitalassetsazure.blob.core.windows.net/rms3-prod/attachments/37/1200x1200/Country-Roasted-Chicken_EXPS_FT20_17623_F_0130_1.jpg"},
		Description: "Roasted chicken with some potatos",
		Time:        "00:40:00",
		Price:       49.99},
	{ID: "6", Name: "Beans", Image: [2]string{
		"https://www.seriouseats.com/thmb/-9Xe6VElxNdWY_MgH2jpuzHIoUA=/450x0/filters:no_upscale():max_bytes(150000):strip_icc()/__opt__aboutcom__coeus__resources__content_migration__serious_eats__seriouseats.com__recipes__images__2014__05__20140530-294420-best-bbq-beans-f286c0d665a548f28ed3dbbff6bc2380.jpg",
		"https://hips.hearstapps.com/hmg-prod.s3.amazonaws.com/images/delish-200318-seo-how-to-cook-beans-horizontal-final-14288-eb-1585337558.jpg?crop=0.6668421052631579xw:1xh;center,top&resize=480:*"},
		Description: "beans",
		Time:        "00:20:00",
		Price:       39.99},
	{ID: "7", Name: "Hot dog", Image: [2]string{
		"https://img.itdg.com.br/tdg/images/recipes/000/160/242/116615/116615_original.jpg?w=1200",
		"https://www.guiadasemana.com.br/contentFiles/image/opt_w1280h960/2021/09/FEA/67630_shutterstock-1150544471.jpg"},
		Description: "hot dog",
		Time:        "00:20:00",
		Price:       39.99},
	{ID: "8", Name: "Sushi", Image: [2]string{
		"https://www.ufrgs.br/laranjanacolher/wp-content/uploads/2021/08/1.png",
		"https://kazasushi.com.br/blog/wp-content/uploads/2018/05/DSC0212-1280x853.jpg"},
		Description: "hot dog",
		Time:        "00:10:00",
		Price:       69.99},
	{ID: "9", Name: "Curry", Image: [2]string{
		"https://images.immediate.co.uk/production/volatile/sites/30/2020/08/113777-0b21d44.jpg?quality=90&resize=400,363",
		"https://assets.bonappetit.com/photos/61a7d632234db3ec67e4fdfa/1:1/w_2000,h_2000,c_limit/20211122%20Japanese%20Curry%20LEDE.jpg"},
		Description: "curry",
		Time:        "01:00:00",
		Price:       59.99},
	{ID: "10", Name: "Soup", Image: [2]string{
		"https://www.inspiredtaste.net/wp-content/uploads/2018/09/Easy-Chicken-Noodle-Soup-Recipe-1200.jpg",
		"https://assets.bonappetit.com/photos/5e31fc9a461d010008e1e976/1:1/w_2560%2Cc_limit/Basically-Lemony-Chicken-Orzo-Soup.jpg"},
		Description: "soup",
		Time:        "01:20:00",
		Price:       39.99},
}

func GetFoods(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, foods)
}

func PostFoods(c *gin.Context) {
	var newFood entities.Food

	if err := c.BindJSON(&newFood); err != nil {
		return
	}

	foods = append(foods, newFood)
	c.IndentedJSON(http.StatusCreated, newFood)
}

func GetFoodByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range foods {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
}

func GetFoodByName(c *gin.Context) {
	name := c.Param("name")

	for _, a := range foods {
		if strings.Contains(strings.ToLower(a.Name), strings.ToLower(name)) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "food not found"})
}
