package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func displayCustomers(c *gin.Context) {
	var customers []Customer
	if err := db.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch customers"})
		return
	}
	c.JSON(http.StatusOK, customers)
}
