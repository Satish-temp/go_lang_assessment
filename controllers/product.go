package controllers

import (
    "net/http"
    "strconv"

    "go-crud-api/models"
    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var prod models.Product
    if err := c.ShouldBindJSON(&prod); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
        return
    }

    prod.ID = models.NextID
    models.NextID++
    models.Products = append(models.Products, prod)

    c.JSON(http.StatusCreated, prod)
}

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, models.Products)
}

func GetProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, p := range models.Products {
        if p.ID == id {
            c.JSON(http.StatusOK, p)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func UpdateProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, p := range models.Products {
        if p.ID == id {
            var updated models.Product
            if err := c.ShouldBindJSON(&updated); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
                return
            }
            updated.ID = id
            models.Products[i] = updated
            c.JSON(http.StatusOK, updated)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func DeleteProduct(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, p := range models.Products {
        if p.ID == id {
            models.Products = append(models.Products[:i], models.Products[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
