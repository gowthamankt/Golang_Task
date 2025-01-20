package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Customer struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name  string    `json:"name"`
	City  string    `json:"city"`
	Email string    `json:"mail_id"`
}

func init() {
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	if err := db.AutoMigrate(&Customer{}); err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	fmt.Println("Connected to the database!")
}

func main() {
	r := gin.Default()

	r.GET("/displayCustomers", displayCustomers)
	r.POST("/createCustomers", Create)
	r.GET("/getCustomerByID/:id", GetCustomerByID)
	r.PUT("/updateCustomer", updateCustomer)
	r.DELETE("/deleteCustomer/:id", deleteCustomer)

	log.Println("Server is running on localhost:5002...")
	if err := r.Run("localhost:5002"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
