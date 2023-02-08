package controllers

import (
	"strconv"

	"stray-dogs/app/models"
	"stray-dogs/pkg/utils"
	"stray-dogs/platform/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
Get the animals using pagination.
Send as a parameter to the endpoint the following:
  - page : int
  - limit : int
*/
func GetAllAminalsByPagination(c *fiber.Ctx) error {
	pageQueryParam := c.Query("page")
	limitQueryParam := c.Query("limit")

	page, err := strconv.Atoi(pageQueryParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send page parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	limit, err := strconv.Atoi(limitQueryParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send limit parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	animals, err := db.AnimalQueries.GetAnimals(page, limit)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "No aminals were found!",
			"count":   0,
			"animals": nil,
		})
	}

	return c.JSON(fiber.Map{
		"count":   len(animals),
		"animals": animals,
	})
}

/*
Create an animal.
*/
func CreateAnimal(c *fiber.Ctx) error {
	id, err := utils.VerifyRoute(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Token not valid",
		})
	}

	animal := &models.Animal{}

	if err := c.BodyParser(animal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if !animal.Chipped {
		animal.ChipNumber = nil
		animal.ChipDate = nil
		animal.ChipPosition = nil
	}

	if animal.Chipped && animal.ChipNumber == nil {
		idChipNumber := uuid.New()
		idChipNumberString := idChipNumber.String()
		animal.ChipNumber = &idChipNumberString
	}

	if !animal.IsAlive && (animal.DeathDate == nil || animal.DeathCause == nil) {
		return c.Status(503).JSON(fiber.Map{
			"error":   true,
			"message": "Please provide death date and death cause!",
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	animalId, err := db.AnimalQueries.CreateAnimalQuery(animal, id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	animal.Id = &animalId

	return c.JSON(fiber.Map{
		"animal": animal,
	})
}

/*
Update and existing animal.
Send as a query paramameter the id of the animal and the new body we want to update.
*/
func UpdateAnimal(c *fiber.Ctx) error {
	userId, err := utils.VerifyRoute(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Token not valid",
		})
	}

	animalQueryId := c.Query("id")

	id, err := strconv.Atoi(animalQueryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send page parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	animal := &models.Animal{}

	if err := c.BodyParser(animal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if !animal.Chipped {
		animal.ChipNumber = nil
		animal.ChipDate = nil
		animal.ChipPosition = nil
	}

	if animal.Chipped && animal.ChipNumber == nil {
		idChipNumber := uuid.New()
		idChipNumberString := idChipNumber.String()
		animal.ChipNumber = &idChipNumberString
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.AnimalQueries.UpdateAnimal(id, animal, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Updated with success!",
		"animal":  animal,
	})
}

/*
Delete an animal.
Send as a query paramameter the id of the animal.
*/
func DeleteAnimal(c *fiber.Ctx) error {
	userId, err := utils.VerifyRoute(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Token not valid",
		})
	}

	animalQueryId := c.Query("id")

	id, err := strconv.Atoi(animalQueryId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send page parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.AnimalQueries.DeleteAnimal(id, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Animal is deleted from the database!",
	})
}

/*
Get animals created by the user.
*/
func GetAnimalsCreatedByUser(c *fiber.Ctx) error {
	userId, err := utils.VerifyRoute(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Token not valid",
		})
	}

	pageQueryParam := c.Query("page")
	limitQueryParam := c.Query("limit")

	page, err := strconv.Atoi(pageQueryParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send page parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	limit, err := strconv.Atoi(limitQueryParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"msg":     "Plase send limit parameter correctly!",
			"count":   0,
			"animals": nil,
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	animals, err := db.AnimalQueries.GetAnimalsCreatedByUser(page, limit, userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "No aminals were found!",
			"count":   0,
			"animals": nil,
		})
	}

	return c.JSON(fiber.Map{
		"count":   len(animals),
		"animals": animals,
	})
}
