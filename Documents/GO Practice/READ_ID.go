package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	var customer Customer
	if err := db.First(&customer, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
