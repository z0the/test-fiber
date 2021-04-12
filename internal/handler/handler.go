package handler

import (
	"encoding/json"
	"strconv"
	"test-job/internal/database"
	"test-job/internal/model"

	"github.com/gofiber/fiber/v2"
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
	app.Get("/organization/list", h.GetOrganizationList)
	app.Post("/organization", h.CreateOrganization)
	app.Put("/organization", h.UpdateOrganization)
	app.Delete("/organization", h.DeleteOrganization)

	return app
}

func (h Handler) CreateOrganization(c *fiber.Ctx) error {
	var org model.Organization
	c.BodyParser(&org)
	newOrg, err := h.db.CreateOrganization(org)
	if err != nil {
		return err
	}
	c.Response().Header.Add("Content-Type", "application/json")
	body, err := json.Marshal(&newOrg)
	if err != nil {
		return err
	}
	c.Status(fiber.StatusCreated)
	return c.Send(body)
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
		return err
	}
	return c.Send(body)
}
func (h Handler) GetOrganizationList(c *fiber.Ctx) error {
	orgs, err := h.db.GetAllOrganizations()
	if err != nil {
		return err
	}
	c.Response().Header.Add("Content-Type", "application/json")
	body, err := json.Marshal(&orgs)
	if err != nil {
		return err
	}
	return c.Send(body)
}

func (h Handler) UpdateOrganization(c *fiber.Ctx) error {
	var org model.Organization
	err := c.BodyParser(&org)
	if err != nil {
		return err
	}
	err = h.db.UpdateOrganization(org)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	c.Response().Header.Add("Content-Type", "application/json")
	return c.Send(c.Body())
}

func (h Handler) DeleteOrganization(c *fiber.Ctx) error {
	orgID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = h.db.DeleteOrganization(orgID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.SendString("Deleted")
}
