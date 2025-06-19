package main

import (
    "go-crud-api/controllers"
    "go-crud-api/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/products", controllers.CreateProduct)
        auth.GET("/products", controllers.GetProducts)
        auth.GET("/products/:id", controllers.GetProduct)
        auth.PUT("/products/:id", controllers.UpdateProduct)
        auth.DELETE("/products/:id", controllers.DeleteProduct)
    }

    r.Run(":8080")
}
