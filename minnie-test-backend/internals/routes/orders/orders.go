package orderRoutes

import (
	"github.com/gofiber/fiber/v2"
	orderHandler "github.com/soberservicesguy/minnie-test-backend/internals/handlers/orders"
)

func SetupOrdersRoutes(router fiber.Router) {
	orders := router.Group("/orders")
	orders.Get("/p/:p/o/:o/s/:start/e/:end/s/:sortType/p/:page/c/:getCount", orderHandler.GetOrders)
}
