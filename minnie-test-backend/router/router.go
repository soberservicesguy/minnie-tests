package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	orderRoutes "github.com/soberservicesguy/minnie-test-backend/internals/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	orderRoutes.SetupOrdersRoutes(api)
}
