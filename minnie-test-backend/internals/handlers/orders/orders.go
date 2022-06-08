package orderHandler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/soberservicesguy/minnie-test-backend/database"
)

// GetNotes func gets all existing notes
// @Description Get all existing notes
// @Tags Notes
// @Accept json
// @Produce json
// @Success 200 {array} model.Note
// @router /api/note [get]
func GetOrders(c *fiber.Ctx) error {

	productName := string(c.Params("p"))
	orderName := string(c.Params("o"))
	startDate := string(c.Params("start"))
	endDate := string(c.Params("end"))
	page, _ := strconv.Atoi(c.Params("page"))
	sortType := string(c.Params("sortType"))
	getCount := string(c.Params("getCount"))

	if orderName == "all" || productName == "all" {
		orderName = ""
		productName = ""
	}

	if sortType == "descend" {
		sortType = "DESC"
	} else {
		sortType = "ASC"
	}
	queryHeader := `
	(
		SELECT order_name, company_name, customer_name, order_time,
			ROUND(SUM(total_amount), 2) AS Total_Amount,
			ROUND(SUM(delivered_amount), 2) AS Delivered_Amount,
			STRING_AGG(product_name, ', ') AS order_products
			FROM (
				SELECT order_name,
					company_name,
					customer_name,
					order_time,
					product_name,
					price,
					quantity,
					delivered_quantity,
					price * quantity AS total_amount,
					price * delivered_quantity AS delivered_amount
					FROM (
						SELECT ordertable.order_name AS order_name, 
							company.company_name AS company_name,
							customer.name AS customer_name,
							ordertable.created_at AS order_time,
							orderitem.product AS product_name,
							orderitem.price_per_unit AS price,
							orderitem.quantity AS quantity,
							delivery.delivered_quantity AS delivered_quantity
							
						FROM ordertable
							JOIN customer ON ordertable.customer_id=customer.user_id
							JOIN company ON customer.company_id=company.company_id
							LEFT JOIN orderitem ON orderitem.order_id=ordertable.id
							LEFT JOIN delivery ON orderitem.id=delivery.order_item_id
					) AS tempQuery 
			) AS finalQuery 
	`
	queryFilter := "WHERE ((product_name LIKE '%" + productName + "%') OR (order_name LIKE '%" + orderName + "%')) AND (finalQuery.order_time BETWEEN DATE '" + startDate + "' AND '" + endDate + "' )"
	groupByCommand := `GROUP BY finalQuery.order_name, finalQuery.company_name, finalQuery.customer_name, finalQuery.order_time`
	sortNOffset := `
			ORDER BY finalQuery.order_time %s
			OFFSET %s LIMIT 5`
	queryCloser := `
		);
	`

	offset := (page - 1) * 5
	fmt.Println(offset)
	queryFooter := fmt.Sprintf(sortNOffset, sortType, strconv.Itoa(offset))
	queryString := queryHeader + queryFilter + groupByCommand + queryFooter + queryCloser
	fmt.Println(queryString)
	db := database.DB
	var orders []map[string]interface{}
	db.Raw(queryString).Scan(&orders)

	if getCount == "true" && len(orders) > 0 {
		queryForCount := queryHeader + queryFilter + groupByCommand + queryCloser
		var ordersForCount []map[string]interface{}
		db.Raw(queryForCount).Scan(&ordersForCount)

		return c.JSON(fiber.Map{"status": "success", "message": "orders Found", "data": orders, "totalRows": len(ordersForCount)})

	} else if getCount == "false" && len(orders) > 0 {

		return c.JSON(fiber.Map{"status": "success", "message": "orders Found", "data": orders})

	}

	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No orders found", "data": nil})

}
