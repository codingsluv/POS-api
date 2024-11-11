package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"

	db "github.com/codingsluv/book-store/POS-api/config"
	"github.com/codingsluv/book-store/POS-api/models"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("registration error: %v", err)
	}
	if data["name"] == "" || data["passCode"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "name and passCode are required",
			"error":   map[string]interface{}{},
		})
	}

	cashier := models.Cashier{
		Name:     data["name"],
		PassCode: data["passCode"],
	}
	db.DB.Create(&cashier)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cashier create error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cashier created",
	})
}

func GetCashier(c *fiber.Ctx) error {
	cashierID := c.Params("cashierID")

	var cashier models.Cashier
	err := db.DB.Where("id =?", cashierID).First(&cashier).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cashier get error",
		})
	}
	cashierData := make(map[string]interface{})
	cashierData["id"] = cashier.ID
	cashierData["name"] = cashier.Name

	if cashier.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "cashier not found",
			"error":   map[string]interface{}{},
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cashier found",
		"data":    cashierData,
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierID := c.Params("cashierID")
	var cashier models.Cashier
	db.DB.Find(&cashier, "id = ?", cashierID)
	if cashier.Name == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "cashier not found",
		})
	}

	var updateCashierData models.Cashier
	c.BodyParser(&updateCashierData)
	if updateCashierData.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "name is required",
			"error":   map[string]interface{}{},
		})
	}
	cashier.Name = updateCashierData.Name
	db.DB.Save(&cashier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cashier updated",
		"data":    cashier,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierID := c.Params("cashierID")
	var cashier models.Cashier
	db.DB.Where("id =?", cashierID).Delete(&cashier)

	if cashier.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "cashier not found",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Where("id = ?", cashierID).Delete(&cashier)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cashier deleted",
	})
}

type Cashier struct {
	ID   uint   `json:"cashierID"`
	Name string `json:"name"`
}

func CashierList(c *fiber.Ctx) error {
	var cashiers []models.Cashier
	db.DB.Find(&cashiers)

	cashierData := make([]map[string]interface{}, len(cashiers))
	for i, cashier := range cashiers {
		cashierData[i] = make(map[string]interface{})
		cashierData[i]["id"] = cashier.ID
		cashierData[i]["name"] = cashier.Name
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "cashiers found",
		"data":    cashierData,
	})
}
