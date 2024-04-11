package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificatonHandler struct {
	store  *db.Store
	object string
}

func NewVerificationHandler(store *db.Store, object string) *VerificatonHandler {
	return &VerificatonHandler{
		store:  store,
		object: object,
	}
}

// Administrative Handler
func (h *VerificatonHandler) HandlePostVerification(c *fiber.Ctx) error {
	var newVerification types.CreateVerificationParams
	if err := c.BodyParser(&newVerification); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	verification, err := newVerification.FromParams()
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: "internal error", Field: "verification failed"})
	}
	verificationCreated, err := h.store.VerificationStore.CreateVerification(c.Context(), verification)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(verificationCreated)
}

func (h *VerificatonHandler) HandleGetVerification(c *fiber.Ctx) error {
	id := c.Params("id")

	verification, err := h.store.VerificationStore.GetVerificationByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(verification)
}

func (h *VerificatonHandler) HandleGetVerifications(c *fiber.Ctx) error {
	verifications, err := h.store.GetVerifications(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(verifications)
}

// Administrative Handler
func (h *VerificatonHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.VerificationStore.DeleteVerification(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "verification not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("User deleted")
}

// Administrative Handler
func (h *VerificatonHandler) HandlePut(c *fiber.Ctx) error {
	var updateParams types.UpdateVerificationParams
	id := c.Params("id")
	if err := c.BodyParser(&updateParams); err != nil {
		return BadRequest(c, err)
	}

	if err := h.store.VerificationStore.UpdateVerification(c.Context(), id, updateParams); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "verification not found", Field: "error"})
		}
		return InternalServerError(c, ErrorObject{Msg: "failed to update verification", Field: "error"})
	}

	return nil
}
