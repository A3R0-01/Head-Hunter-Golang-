package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyHandler struct {
	store  *db.Store
	object string
}

func NewCompanyHandler(store *db.Store, object string) *CompanyHandler {
	return &CompanyHandler{
		store:  store,
		object: object,
	}
}

func (h *CompanyHandler) HandlePostCompany(c *fiber.Ctx) error {
	var newCompany types.CreateCompanyParams
	if err := c.BodyParser(&newCompany); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	if errors := newCompany.Validate(); len(errors) > 0 {
		return BadRequest(c, errors)
	}
	hr, err := newCompany.HeadOfHr.FromParams()
	if err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid user/hr", Field: "error"})
	}
	chanUser := make(chan *types.User, 1)
	chanError := make(chan error, 1)
	go func(userChan chan *types.User, errorChan chan error, hr *types.User) {
		user, err := h.store.UserStore.CreateUser(c.Context(), hr)
		if err != nil {
			errorChan <- err
		}
		userChan <- user
	}(chanUser, chanError, hr)
	company := newCompany.FromParams()
	select {
	case user := <-chanUser:
		company.HeadOfHr = user.ID
	case err := <-chanError:
		return InternalServerError(c, err)
	}
	companyCreated, err := h.store.CompanyStore.CreateCompany(c.Context(), company)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(companyCreated)
}

func (h *CompanyHandler) HandleGetCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	company, err := h.store.CompanyStore.GetCompanyByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(company)
}

func (h *CompanyHandler) HandleGetCompanies(c *fiber.Ctx) error {
	companies, err := h.store.GetCompanies(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(companies)
}

func (h *CompanyHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.CompanyStore.DeleteCompany(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "company not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Company deleted")
}

func (h *CompanyHandler) HandlePut(c *fiber.Ctx) error {
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
