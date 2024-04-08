package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndustryHandler struct {
	store  *db.Store
	object string
}

func NewIndustryHandler(store *db.Store, object string) *IndustryHandler {
	return &IndustryHandler{
		store:  store,
		object: object,
	}
}

func (h *IndustryHandler) HandlePostIndustry(c *fiber.Ctx) error {
	var newIndustry types.CreateIndustryParams
	if err := c.BodyParser(&newIndustry); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	industryFromParams, err := newIndustry.FromParams()
	if err != nil {
		return InternalServerError(c, ErrorObject{Msg: "we hit a snag", Field: "error"})
	}
	industryCreated, err := h.store.IndustryStore.CreateIndustry(c.Context(), industryFromParams)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(industryCreated)
}

func (h *IndustryHandler) HandleGetIndustry(c *fiber.Ctx) error {
	id := c.Params("id")

	industry, err := h.store.IndustryStore.GetIndustryByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(industry)
}

func (h *IndustryHandler) HandleGetIndustries(c *fiber.Ctx) error {
	companies, err := h.store.GetCompanies(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(companies)
}

func (h *IndustryHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.CompanyStore.DeleteCompany(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "company not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Company deleted")
}

func (h *IndustryHandler) HandlePut(c *fiber.Ctx) error {
	var updateParams types.UpdateCompanyParams
	id := c.Params("id")
	if err := c.BodyParser(&updateParams); err != nil {
		return BadRequest(c, err)
	}

	if err := h.store.CompanyStore.UpdateCompany(c.Context(), id, updateParams); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "company not found", Field: "error"})
		}
		return InternalServerError(c, ErrorObject{Msg: "failed to update company", Field: "error"})
	}

	return nil
}
