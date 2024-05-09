package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionHandler struct {
	store  *db.Store
	object string
}

func NewSessionHandler(store *db.Store, object string) *SessionHandler {
	return &SessionHandler{
		store:  store,
		object: object,
	}
}

func (h *SessionHandler) HandlePostSession(c *fiber.Ctx) error {
	var newSession types.CreateSessionParams
	if err := c.BodyParser(&newSession); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	session, err := newSession.FromParams()
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: "error with the parameters", Field: "session failed"})
	}
	sessionCreated, err := h.store.SessionStore.CreateSession(c.Context(), session)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(sessionCreated)
}

func (h *SessionHandler) HandleGetSession(c *fiber.Ctx) error {
	id := c.Params("id")

	session, err := h.store.SessionStore.GetSessionByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(session)
}

func (h *SessionHandler) HandleGetSessions(c *fiber.Ctx) error {
	sessions, err := h.store.GetSessions(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(sessions)
}

func (h *SessionHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.SessionStore.DeleteSession(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "session not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Session deleted")
}

func (h *SessionHandler) HandlePut(c *fiber.Ctx) error {
	var updateParams types.UpdateSessionParams
	id := c.Params("id")
	if err := c.BodyParser(&updateParams); err != nil {
		return BadRequest(c, err)
	}

	if err := h.store.SessionStore.UpdateSession(c.Context(), id, &updateParams); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "session not found", Field: "error"})
		}
		return InternalServerError(c, ErrorObject{Msg: "failed to update session", Field: "error"})
	}

	return nil
}
