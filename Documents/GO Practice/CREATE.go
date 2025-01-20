package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func Create(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = uuid.New()

	if err := db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create customer"})
		return
	}

	c.JSON(http.StatusOK, customer)
}
