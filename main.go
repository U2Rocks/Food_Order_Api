package main

import (
	"Food_Order_Api/database"
	"Food_Order_Api/models"
	"Food_Order_Api/routes"
	"log"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// function to initialize routes on api startup
func initialize_routes(app *fiber.App) {
	// get routes
	app.Get("/orders/list", routes.ListAllOrders)                 // works
	app.Get("/orders/incomplete", routes.ListAllIncompleteOrders) // works
	app.Get("/orders/complete", routes.ListAllCompleteOrders)     // works

	// post routes
	app.Post("/orders/add", routes.AddNewOrder)          // works
	app.Post("/orders/completion", routes.CompleteOrder) // works
	app.Post("/orders/update", routes.UpdateOrder)       // works
	app.Post("/orders/current", routes.CurrentOrder)     // works

	// delete routes
	app.Post("/orders/delete", routes.DeleteOrder) // works

}

// function that sets up database connection and migrates tables
func initialize_database() {
	var err error
	database.DBC, err = gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	fmt.Println("Connection to Database Now Open")
	database.DBC.AutoMigrate(&models.FoodOrder{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()     // create a new fiber app instance
	initialize_database()  // open database connection and migrate tables
	initialize_routes(app) // create all routes

	// static route for root that shows a nice page
	app.Static("/", "./public/index.html", fiber.Static{Index: "index.html"})

	// app.Listen wrapped in fatal log call
	log.Fatal(app.Listen(":3000"))
}
