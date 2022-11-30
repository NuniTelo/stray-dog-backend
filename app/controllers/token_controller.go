package controllers

import (
	"stray-dogs/app/models"
	"stray-dogs/platform/database"

	"github.com/gofiber/fiber/v2"
)

func LoginAndGetNewToken(c *fiber.Ctx) error {
	userLoginStruct := &models.UserLoginStruct{}

	if err := c.BodyParser(userLoginStruct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token, err := db.UserQueries.Login(userLoginStruct)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"token": token,
	})
}

func CreateNewAdminUser(c *fiber.Ctx) error {
	userBody := &models.User{}

	if err := c.BodyParser(userBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.UserQueries.CreateAdminUser(userBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Success",
	})
}
