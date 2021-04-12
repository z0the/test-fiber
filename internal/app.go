package internal

import (
	"test-job/internal/database"
	"test-job/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	db := database.NewPostgresDB()
	h := handler.NewHandler(db)
	app := h.InitRoutes(fiber.Config{})

	app.Listen(":3000")
}
