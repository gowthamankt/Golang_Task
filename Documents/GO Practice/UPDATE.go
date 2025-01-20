package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
