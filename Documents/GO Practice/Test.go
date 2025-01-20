package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	r := gin.Default()
	r.POST("/createCustomers", Create)

	customer := Customer{
		Name:  "John Doe",
		City:  "New York",
		Email: "john.doe@example.com",
	}
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/createCustomers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var createdCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&createdCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.Name, createdCustomer.Name)
	assert.Equal(t, customer.City, createdCustomer.City)
	assert.Equal(t, customer.Email, createdCustomer.Email)
	assert.NotEqual(t, uuid.Nil, createdCustomer.ID) // Check if ID is generated
}

func TestGetCustomerByID(t *testing.T) {
	r := gin.Default()
	r.GET("/getCustomerByID/:id", GetCustomerByID)

	customer := Customer{
		Name:  "Jane Doe",
		City:  "Los Angeles",
		Email: "jane.doe@example.com",
	}
	db.Create(&customer)

	req, err := http.NewRequest("GET", "/getCustomerByID/"+customer.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var retrievedCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&retrievedCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, retrievedCustomer.ID)
}

func TestDisplayCustomers(t *testing.T) {
	r := gin.Default()
	r.GET("/displayCustomers", displayCustomers)

	customer := Customer{
		Name:  "Alice",
		City:  "San Francisco",
		Email: "alice@example.com",
	}
	db.Create(&customer)

	req, err := http.NewRequest("GET", "/displayCustomers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var customers []Customer
	err = json.NewDecoder(rr.Body).Decode(&customers)
	assert.Nil(t, err)
	assert.True(t, len(customers) > 0)
}

func TestUpdateCustomer(t *testing.T) {
	r := gin.Default()
	r.POST("/updateCustomer", updateCustomer)

	customer := Customer{
		Name:  "Bob",
		City:  "Chicago",
		Email: "bob@example.com",
	}
	db.Create(&customer)

	customer.Name = "Robert"
	customer.City = "Houston"
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/updateCustomer", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var updatedCustomer Customer
	err = json.NewDecoder(rr.Body).Decode(&updatedCustomer)
	assert.Nil(t, err)
	assert.Equal(t, customer.Name, updatedCustomer.Name)
	assert.Equal(t, customer.City, updatedCustomer.City)
}

func TestDeleteCustomer(t *testing.T) {
	r := gin.Default()
	r.DELETE("/deleteCustomer/:id", deleteCustomer)

	customer := Customer{
		Name:  "Eve",
		City:  "Miami",
		Email: "eve@example.com",
	}
	db.Create(&customer)

	req, err := http.NewRequest("DELETE", "/deleteCustomer/"+customer.ID.String(), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code) // Assuming success returns 200, or you can choose to return 204 if no content is returned
}
