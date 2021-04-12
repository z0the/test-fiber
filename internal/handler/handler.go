package handler

import (
	"encoding/json"
	"strconv"
	"test-job/internal/database"
	"test-job/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	db database.Database
}

func NewHandler(db database.Database) *Handler {
	return &Handler{db: db}
}
func (h Handler) InitRoutes(config fiber.Config) *fiber.App {
	app := fiber.New(config)

	app.Get("/organization", h.GetOrganization)
	app.Post("/organization", h.CreateOrganization)
	return app
}

func (h Handler) CreateOrganization(c *fiber.Ctx) error {
	var org model.Organization
	c.BodyParser(&org)
	err := h.db.CreateOrganization(org)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	logrus.Info("OrgName:", org.Name)
	c.Status(fiber.StatusCreated)
	return c.SendString("Created")
}

func (h Handler) GetOrganization(c *fiber.Ctx) error {
	orgID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	org, err := h.db.GetOrganization(orgID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	c.Response().Header.Add("Content-Type", "application/json")
	body, err := json.Marshal(&org)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Send(body)
}
