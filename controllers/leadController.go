package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudheerpal/learn-go/database"
	"github.com/sudheerpal/learn-go/models"
)

func AllLeads(c *fiber.Ctx) error {

	var leads []models.Lead

	database.DB.Find(&leads)

	return c.JSON(leads)

}
