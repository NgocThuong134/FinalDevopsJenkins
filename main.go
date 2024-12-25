package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: "1", Name: "Item 1"},
    {ID: "2", Name: "Item 2"},
}

func main() {
    router := gin.Default()

	router.GET("/ping", ping)
	router.GET("/hello", hello)

    router.GET("/items", getItems)
    router.GET("/items/:id", getItemByID)
    router.POST("/items", createItem)
    router.PUT("/items/:id", updateItem)
    router.DELETE("/items/:id", deleteItem)

    router.Run(":7080")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world !!!"})
}

func ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

func getItemByID(c *gin.Context) {
    id := c.Param("id")
    for _, item := range items {
        if item.ID == id {
            c.JSON(http.StatusOK, item)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func createItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err == nil {
        items = append(items, newItem)
        c.JSON(http.StatusCreated, newItem)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func updateItem(c *gin.Context) {
    id := c.Param("id")
    var updatedItem Item
    if err := c.ShouldBindJSON(&updatedItem); err == nil {
        for i, item := range items {
            if item.ID == id {
                items[i].Name = updatedItem.Name
                c.JSON(http.StatusOK, items[i])
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func deleteItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range items {
        if item.ID == id {
            items = append(items[:i], items[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}