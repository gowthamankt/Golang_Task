package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var input struct {
			ID uuid.UUID `json:"id"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		var customer Customer

		result := db.First(&customer, "id = ?", input.ID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Customer not found", http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("Error fetching customer: %v", result.Error), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
