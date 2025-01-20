package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteCustomer(c *gin.Context) {
	id := c.Param("id")

	if err := db.Delete(&Customer{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
