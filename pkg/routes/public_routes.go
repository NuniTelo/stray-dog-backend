package routes

import (
	"stray-dogs/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/animals", controllers.GetAllAminalsByPagination)
	route.Post("/login", controllers.LoginAndGetNewToken)
	route.Post("/admin/create", controllers.CreateNewAdminUser)
}
