package orderRoutes

import (
	"github.com/gmvbr/httptest"
	"github.com/gofiber/fiber/v2"
	"github.com/soberservicesguy/minnie-test-backend/database"
	"testing"

	orderHandler "github.com/soberservicesguy/minnie-test-backend/internals/handlers/orders"
)

func TestHandler(t *testing.T) {
	app := fiber.New()
	database.ConnectDB("../../.env")
	app.Get("/p/:p/o/:o/s/:start/e/:end/s/:sortType/p/:page/c/:getCount", func(c *fiber.Ctx) error {
		return orderHandler.GetOrders(c)
	})

	// make query which returns data, page count, in descending order and status code 200
	test1 := httptest.Get("/p/ani/o/100/s/2020-01-01/e/2020-02-02/s/descend/p/1/c/true").Test(t, app)
	test1.Status(200)
	test1.Body("{\"data\":[{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #002-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-15T22:34:12+05:00\",\"total_amount\":null},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":3198,\"order_name\":\"PO #002-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-15T22:34:12+05:00\",\"total_amount\":3198},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #003-I\",\"order_products\":\"Hand sanitiZER\",\"order_time\":\"2020-01-05T10:34:12+05:00\",\"total_amount\":6008.71},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":4674.33,\"order_name\":\"PO #003-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-05T10:34:12+05:00\",\"total_amount\":4674.33},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #005-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-03T15:34:12+05:00\",\"total_amount\":299.21}],\"message\":\"orders Found\",\"status\":\"success\",\"totalRows\":8}")

	// make query which returns data, page count, in ascending order and status code 200
	test2 := httptest.Get("/p/ani/o/100/s/2020-01-01/e/2020-02-02/s/ascend/p/1/c/true").Test(t, app)
	test2.Status(200)
	test2.Body("{\"data\":[{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #001-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-02T20:34:12+05:00\",\"total_amount\":904.67},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":275,\"order_name\":\"PO #001-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-02T20:34:12+05:00\",\"total_amount\":275},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #005-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-03T15:34:12+05:00\",\"total_amount\":299.21},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":3480,\"order_name\":\"PO #005-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-03T15:34:12+05:00\",\"total_amount\":3480},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #003-I\",\"order_products\":\"Hand sanitiZER\",\"order_time\":\"2020-01-05T10:34:12+05:00\",\"total_amount\":6008.71}],\"message\":\"orders Found\",\"status\":\"success\",\"totalRows\":8}")

	// make query which returns data, no page count and status code 200
	test3 := httptest.Get("/p/ani/o/100/s/2020-01-01/e/2020-02-02/s/descend/p/1/c/false").Test(t, app)
	test3.Status(200)
	test3.Body("{\"data\":[{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #002-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-15T22:34:12+05:00\",\"total_amount\":null},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":3198,\"order_name\":\"PO #002-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-15T22:34:12+05:00\",\"total_amount\":3198},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #003-I\",\"order_products\":\"Hand sanitiZER\",\"order_time\":\"2020-01-05T10:34:12+05:00\",\"total_amount\":6008.71},{\"company_name\":\"Pupkin \\u0026 Co\",\"customer_name\":\"Petr Petrovich\",\"delivered_amount\":4674.33,\"order_name\":\"PO #003-P\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-05T10:34:12+05:00\",\"total_amount\":4674.33},{\"company_name\":\"Roga \\u0026 Kopyta\",\"customer_name\":\"Ivan Ivanovich\",\"delivered_amount\":null,\"order_name\":\"PO #005-I\",\"order_products\":\"Hand sanitizer\",\"order_time\":\"2020-01-03T15:34:12+05:00\",\"total_amount\":299.21}],\"message\":\"orders Found\",\"status\":\"success\"}")

	// make query which doesnt return data and status code 404
	test4 := httptest.Get("/p/anis/o/100/s/2020-01-01/e/2020-02-02/s/descend/p/1/c/true").Test(t, app)
	test4.Status(404)
	test4.Body("{\"data\":null,\"message\":\"No orders found\",\"status\":\"error\"}")

}
