package users

import (
	"log"
	"market/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func List(c *gin.Context) {
	var users []User

	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"result": users,
	})
}

func Get(c *gin.Context) {
	id := c.Param("id")

	var user User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	c.Bind(&body)

	var user User
	result := initializers.DB.Where(&User{Username: body.Username}).First(&user)

	if result.Error == nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": "Done.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "Not matched.",
		})
	}
}

func Create(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	c.Bind(&body)

	encrypt, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Fail to encrypt password")
		return
	}

	user := User{Username: body.Username, Password: string(encrypt[:])}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func Modify(c *gin.Context) {
	var body struct {
		ID       uint
		Username string
		Password string
	}
	c.Bind(&body)

	var user User
	result := initializers.DB.First(&user, body.ID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	result = initializers.DB.Model(&user).Updates(body)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": body.ID,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&User{}, id)

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
