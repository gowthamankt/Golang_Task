package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Create customers
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var customers []Customer

		// Decode the JSON request body into the customers slice
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&customers)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Insert each customer into the database
		for _, customer := range customers {
			// Generate a new UUID if it's not provided
			if customer.ID == uuid.Nil {
				customer.ID = uuid.New()
			}

			// Insert the customer into the database using GORM
			if err := db.Create(&customer).Error; err != nil {
				http.Error(w, fmt.Sprintf("Error inserting customer: %v", err), http.StatusInternalServerError)
				return
			}
		}

		// Send a success message
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Customers added successfully!"})

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
