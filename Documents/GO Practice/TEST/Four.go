package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Get a customer by ID
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var input struct {
			ID uuid.UUID `json:"id"`
		}

		// Decode the input JSON containing the customer ID
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		var customer Customer

		// Retrieve the customer by ID using GORM's First() method
		result := db.First(&customer, "id = ?", input.ID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Customer not found", http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Error fetching customer: %v", result.Error), http.StatusInternalServerError)
			}
			return
		}

		// Send the customer as JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
