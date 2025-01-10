package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Update an existing customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var customer Customer

		// Decode the JSON request body into the customer struct
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&customer)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Check if customer ID is valid
		if customer.ID == uuid.Nil {
			http.Error(w, "Customer ID is required", http.StatusBadRequest)
			return
		}

		// Find the customer by ID
		var existingCustomer Customer
		result := db.First(&existingCustomer, "id = ?", customer.ID)
		if result.Error != nil {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}

		// Update the customer fields
		existingCustomer.Name = customer.Name
		existingCustomer.City = customer.City
		existingCustomer.Email = customer.Email

		// Save the updated customer
		if err := db.Save(&existingCustomer).Error; err != nil {
			http.Error(w, fmt.Sprintf("Error updating customer: %v", err), http.StatusInternalServerError)
			return
		}

		// Send a success message
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Customer updated successfully!"})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
