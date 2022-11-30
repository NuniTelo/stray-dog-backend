package routes

import (
	"stray-dogs/app/controllers"

	"stray-dogs/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Put("/animal", middleware.JWTProtected(), controllers.CreateAnimal)
	route.Get("/profile/animals", middleware.JWTProtected(), controllers.GetAnimalsCreatedByUser)
	route.Post("/animal", middleware.JWTProtected(), controllers.UpdateAnimal)
	route.Delete("/animal", middleware.JWTProtected(), controllers.DeleteAnimal)

	route.Post("/image/upload", middleware.JWTProtected(), controllers.FileUpload)
}
