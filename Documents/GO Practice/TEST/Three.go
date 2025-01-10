package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Read all customers
func displayCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var customers []Customer

		// Retrieve all customers from the database using GORM
		if err := db.Find(&customers).Error; err != nil {
			http.Error(w, fmt.Sprintf("Error fetching customers: %v", err), http.StatusInternalServerError)
			return
		}

		// Send the customers as JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
