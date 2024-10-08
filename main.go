package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/routes"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	database.Connect()

	app := fiber.New()

	// Initialize default config
	app.Use(logger.New())

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://mspos.onrender.com, http://localhost:4200",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(getPort()))

}
