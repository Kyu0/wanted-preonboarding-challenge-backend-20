package products

import (
	"market/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var products []Product

	initializers.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"result": products,
	})
}

func Get(c *gin.Context) {
	id := c.Param("id")

	var product Product
	result := initializers.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}

func Create(c *gin.Context) {
	var body struct {
		Name   string
		Price  uint32
		Status TransactionStatus `binding:"enum"`
		Amount uint16
		UserId uint
	}
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	product := Product{Name: body.Name, Price: body.Price, Status: body.Status, Amount: body.Amount, UserId: body.UserId}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}

func Modify(c *gin.Context) {
	var body struct {
		ID     uint
		Name   string
		Price  uint32
		Status TransactionStatus
		Amount uint16
		UserId uint
	}
	c.Bind(&body)

	var product Product
	result := initializers.DB.First(&product, body.ID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	result = initializers.DB.Model(&product).Updates(body)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": body,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&Product{}, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "No records have that id.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": id,
	})
}
