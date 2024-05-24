package main

import (
	"market/initializers"
	"market/products"
	"market/users"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.Connect()
	initializers.DB.AutoMigrate(&users.User{}, &products.Product{})
}

func setUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", users.List)
	r.GET("/users/:id", users.Get)
	r.POST("/users", users.Create)
	r.POST("/login", users.Login)
	r.PUT("/users", users.Modify)
	r.DELETE("/users/:id", users.Delete)

	return r
}

func main() {
	r := setUpRouter()
	r.Run()
}
