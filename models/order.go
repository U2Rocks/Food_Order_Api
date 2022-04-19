package models

import (
	"gorm.io/gorm"
)

// struct used to create database and models a whole food order
type FoodOrder struct {
	gorm.Model
	Cfirst            string  `json:"cfirst"`
	Clast             string  `json:"clast"`
	Items             string  `json:"items"`
	Description       string  `json:"description"`
	Total_Cost        float32 `json:"total_cost"`
	Completion_Status bool    `json:"completion_status"`
}

// struct used to update database fields like Items and Description
type OrderUpdate struct {
	Cfirst      string  `json:"cfirst"`
	Clast       string  `json:"clast"`
	Items       string  `json:"items"`
	Description string  `json:"description"`
	Total_Cost  float32 `json:"total_cost"`
}

// struct used to select DB row to soft delete
type DeleteOrder struct {
	Cfirst            string `json:"cfirst"`
	Clast             string `json:"clast"`
	Completion_Status bool   `json:"completion_status"`
}

// struct used to complete an order so that it does not show up in queries
type CompleteOrder struct {
	Cfirst            string `json:"cfirst"`
	Clast             string `json:"clast"`
	Completion_Status bool   `json:"completion_status"`
}

// struct used to fetch current order information for a customer
type CurrentOrder struct {
	Cfirst string `json:"cfirst"`
	Clast  string `json:"clast"`
}
