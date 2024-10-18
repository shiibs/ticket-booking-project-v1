package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shiibs/ticket-booking-project-v1/config"
	"github.com/shiibs/ticket-booking-project-v1/db"
	"github.com/shiibs/ticket-booking-project-v1/handlers"
	"github.com/shiibs/ticket-booking-project-v1/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	server := app.Group("/api")

	eventRepository := repositories.NewEventRepository(db)

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))

}
