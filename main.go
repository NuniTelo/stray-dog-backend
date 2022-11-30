package main

import (
	"stray-dogs/pkg/configs"
	"stray-dogs/pkg/middleware"
	"stray-dogs/pkg/routes"
	"stray-dogs/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServer(app)
}
