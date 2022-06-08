package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/soberservicesguy/minnie-test-backend/database"
	"github.com/soberservicesguy/minnie-test-backend/router"
)

func main() {
	app := fiber.New()
	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDB("./.env")
	pwd, _ := os.Getwd()
	fmt.Println("")
	fmt.Println(pwd)
	fmt.Println("")
	app.Static("/", pwd+"/frontend")
	app.Static("/home", pwd+"/frontend")
	app.Static("/orders", pwd+"/frontend")
	router.SetupRoutes(app)
	app.Listen(":3000")
}
