package main

import (
	"market/initializers"
	"market/users"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()

	r.GET("/users", users.List)
	r.GET("/users/:id", users.Get)

	r.Run()
}
