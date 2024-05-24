package main

import (
	"market/initializers"
	"market/products"
	"market/users"
	"market/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.Connect()
	initializers.DB.AutoMigrate(&users.User{}, &products.Product{})
}

func ValidateEnum(level validator.FieldLevel) bool {
	value := level.Field().Interface().(util.Enum)

	return value.IsVaild()
}

func setUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", users.List)
	r.GET("/users/:id", users.Get)
	r.POST("/users", users.Create)
	r.POST("/login", users.Login)
	r.PUT("/users", users.Modify)
	r.DELETE("/users/:id", users.Delete)

	r.GET("/products", products.List)
	r.GET("/products/:id", products.Get)
	r.POST("/products", products.Create)
	r.PUT("/products", products.Modify)
	r.DELETE("/products/:id", products.Delete)

	return r
}

func main() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("enum", ValidateEnum)
	}

	r := setUpRouter()
	r.Run()
}
