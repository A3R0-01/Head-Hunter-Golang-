package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationHandler struct {
	store  *db.Store
	object string
}

func NewApplicationHandler(store *db.Store, object string) *ApplicationHandler {
	return &ApplicationHandler{
		store:  store,
		object: object,
	}
}

func (h *ApplicationHandler) HandlePostApplication(c *fiber.Ctx) error {
	var newApplication types.CreateApplicationParams
	if err := c.BodyParser(&newApplication); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	application, err := newApplication.FromParams()
	if err != nil {
		return InternalServerError(c, ErrorObject{Msg: "internal error", Field: "Application failed"})
	}
	applicationCreated, err := h.store.ApplicationStore.CreateApplication(c.Context(), application)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(applicationCreated)
}

func (h *ApplicationHandler) HandleGetApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	application, err := h.store.ApplicationStore.GetApplicationByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(application)
}

func (h *ApplicationHandler) HandleGetApplications(c *fiber.Ctx) error {
	applications, err := h.store.GetApplications(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(applications)
}

func (h *ApplicationHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.ApplicationStore.DeleteApplication(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "Application not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Application deleted")
}
