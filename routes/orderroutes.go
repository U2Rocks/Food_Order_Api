package routes

import (
	"Food_Order_Api/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	//"gorm.io/gorm"
	"Food_Order_Api/database"
)

// function that creates a new order in the database
func AddNewOrder(c *fiber.Ctx) error {
	db := database.DBC
	order := new(models.FoodOrder)

	if err := c.BodyParser(order); err != nil {
		Response := models.ResponseObject{Statuscode: 500, Message: "Could not parse Request"}
		return c.Status(500).JSON(Response)
	}

	db.Create(&order)
	PositiveResponse := models.ResponseObject{Statuscode: 201, Message: "Successfully created order"}
	return c.Status(201).JSON(PositiveResponse)
}

// function that lists all orders in the database
func ListAllOrders(c *fiber.Ctx) error {
	db := database.DBC
	var orderList []models.FoodOrder

	db.Find(&orderList)
	return c.JSON(orderList)
}

// function that lists all incomplete orders in the database(not tested)
func ListAllIncompleteOrders(c *fiber.Ctx) error {
	db := database.DBC
	var orderList []models.FoodOrder

	db.Raw("SELECT * FROM food_orders WHERE completion_status = false AND deleted_at IS NULL").Scan(&orderList)
	return c.JSON(orderList)
}

// function that lists all complete orders in the database(not tested)
func ListAllCompleteOrders(c *fiber.Ctx) error {
	db := database.DBC
	var orderList []models.FoodOrder

	db.Raw("SELECT * FROM food_orders WHERE completion_status = true AND deleted_at IS NULL").Scan(&orderList)
	return c.JSON(orderList)
}

// function that deletes an order from the database
func DeleteOrder(c *fiber.Ctx) error {
	db := database.DBC
	deleteObject := new(models.DeleteOrder)
	var order models.FoodOrder

	if err := c.BodyParser(deleteObject); err != nil {
		Response := models.ResponseObject{Statuscode: 500, Message: "Could not parse delete request body"}
		return c.Status(500).JSON(Response)
	}

	stringBool := strconv.FormatBool(deleteObject.Completion_Status)
	correctedBool := strings.ToUpper(stringBool)

	db.Raw("SELECT * FROM food_orders WHERE cfirst = '" + deleteObject.Cfirst + "' AND clast = '" + deleteObject.Clast + "' AND completion_status = " + correctedBool).Scan(&order)

	if order.ID == 0 {
		BadResponse := models.ResponseObject{Statuscode: 500, Message: "Malformed query"}
		return c.Status(500).JSON(BadResponse)
	}

	db.Delete(&models.FoodOrder{}, order.ID)
	PostiveResponse := models.ResponseObject{Statuscode: 204, Message: "Entry Successfully Deleted"}
	return c.Status(202).JSON(PostiveResponse)
}

// function that marks an order as complete
func CompleteOrder(c *fiber.Ctx) error {
	db := database.DBC
	orderData := new(models.CompleteOrder)
	var order models.FoodOrder

	if err := c.BodyParser(orderData); err != nil {
		Response := models.ResponseObject{Statuscode: 500, Message: "Could not parse complete request body"}
		return c.Status(500).JSON(Response)
	}

	db.Raw("SELECT * FROM food_orders WHERE cfirst = '" + orderData.Cfirst + "' AND clast = '" + orderData.Clast + "' AND completion_status = false AND deleted_at IS NULL").Scan(&order)
	order.Completion_Status = true
	db.Save(&order)

	PositiveResponse := models.ResponseObject{Statuscode: 200, Message: "Order Successfully Marked as complete"}
	return c.Status(200).JSON(PositiveResponse)
}

// function that updates an order with new information(can update items/cost -> description -> or both)
func UpdateOrder(c *fiber.Ctx) error {
	db := database.DBC
	updateData := new(models.OrderUpdate)
	var order models.FoodOrder

	if err := c.BodyParser(updateData); err != nil {
		Response := models.ResponseObject{Statuscode: 500, Message: "Could not parse update body request"}
		return c.Status(500).JSON(Response)
	}

	// query the database and scan appropriate entry(make new entry...)
	db.Raw("SELECT * FROM food_orders WHERE cfirst = '" + updateData.Cfirst + "' AND clast = '" + updateData.Clast + "' AND completion_status = false AND deleted_at IS NULL").Scan(&order)

	// check if orderid is 0 and report that query was malformed
	if order.ID == 0 {
		BadResponse := models.ResponseObject{Statuscode: 500, Message: "Malformed query"}
		return c.Status(500).JSON(BadResponse)
	}

	// update a part of the record and save to db
	if updateData.Items == "" {
		order.Description = updateData.Description
		db.Save(&order)
	} else if updateData.Description == "" {
		order.Items = updateData.Items
		order.Total_Cost = updateData.Total_Cost
		db.Save(&order)
	} else {
		order.Items = updateData.Items
		order.Description = updateData.Description
		order.Total_Cost = updateData.Total_Cost
		db.Save(&order)
	}

	PositiveResponse := models.ResponseObject{Statuscode: 204, Message: "Order successfully updated"}
	return c.Status(204).JSON(PositiveResponse)
}

// function to return the current order for a given name
func CurrentOrder(c *fiber.Ctx) error {
	db := database.DBC
	currentData := new(models.CurrentOrder)
	var order models.FoodOrder

	if err := c.BodyParser(currentData); err != nil {
		Response := models.ResponseObject{Statuscode: 500, Message: "Could not parse current order request body"}
		return c.Status(500).JSON(Response)
	}

	db.Raw("SELECT * FROM food_orders Where cfirst = '" + currentData.Cfirst + "' AND clast = '" + currentData.Clast + "' AND completion_status = false AND deleted_at IS NULL").Scan(&order)

	// check if orderid is 0 and report that query was malformed
	if order.ID == 0 {
		BadResponse := models.ResponseObject{Statuscode: 500, Message: "Malformed query"}
		return c.Status(500).JSON(BadResponse)
	}

	return c.Status(200).JSON(order)
}
