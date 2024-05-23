package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

}

func Get(c *gin.Context) {
	var user User

	c.JSON(http.StatusOK, gin.H{
		"users": user,
	})
}

func Create(c *gin.Context) {

}

func Modify(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
