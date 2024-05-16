package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageHandler struct {
	store  *db.Store
	object string
}

func NewMessageHandler(store *db.Store, object string) *MessageHandler {
	return &MessageHandler{
		store:  store,
		object: object,
	}
}

func (h *MessageHandler) HandlePostMessage(c *fiber.Ctx) error {
	var newMessage types.CreateMessageParams
	if err := c.BodyParser(&newMessage); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	message, err := newMessage.FromParams()
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: "Message failed", Field: "an error occured on your end"})
	}
	messageCreated, err := h.store.MessageStore.CreateMessage(c.Context(), message)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(messageCreated)
}

func (h *MessageHandler) HandleGetMessage(c *fiber.Ctx) error {
	id := c.Params("id")

	message, err := h.store.MessageStore.GetMessageByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(message)
}

func (h *MessageHandler) HandleGetMessages(c *fiber.Ctx) error {
	messages, err := h.store.GetMessages(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(messages)
}

func (h *MessageHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.MessageStore.DeleteMessage(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "Message not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Message deleted")
}
