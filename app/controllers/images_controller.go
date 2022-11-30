package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"stray-dogs/app/models"
	"stray-dogs/app/queries"
	"stray-dogs/pkg/utils"
)

func FileUpload(c *fiber.Ctx) error {
	_, err := utils.VerifyRoute(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Token not valid",
		})
	}

	formHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"StatusCode": http.StatusInternalServerError,
				"Message":    "error",
				"Data":       "Select a file to upload",
			})
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"StatusCode": http.StatusInternalServerError,
				"Message":    "error",
				"Data":       err.Error(),
			})
	}
	mediaStructReference := &queries.Media{}

	uploadUrl, err := mediaStructReference.FileUpload(models.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{
				"StatusCode": http.StatusInternalServerError,
				"Message":    "error",
				"Data":       err.Error(),
			})
	}

	return c.Status(http.StatusOK).JSON(
		fiber.Map{
			"StatusCode": http.StatusInternalServerError,
			"Message":    "Uploaded successfully",
			"Data":       uploadUrl,
		})
}
