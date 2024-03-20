package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	store  *db.Store
	object string
}

func NewUserHandler(store *db.Store, object string) *UserHandler {
	return &UserHandler{
		store:  store,
		object: object,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var newUser types.CreateUserParams
	if err := c.BodyParser(&newUser); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	if errors := newUser.Validate(); len(errors) > 0 {
		return BadRequest(c, errors)
	}
	user, err := newUser.FromParams()
	if err != nil {
		return InternalServerError(c, ErrorObject{Msg: "internal error", Field: "user failed"})
	}
	userCreated, err := h.store.UserStore.CreateUser(c.Context(), user)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(userCreated)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.store.UserStore.GetUserByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.store.GetUsers(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(users)
}

func (h *UserHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.UserStore.DeleteUser(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "user not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("User deleted")
}

func (h *UserHandler) HandlePut(c *fiber.Ctx) error {
	var updateParams types.UpdateUserParams
	id := c.Params("id")
	if err := c.BodyParser(&updateParams); err != nil {
		return BadRequest(c, err)
	}

	if err := h.store.UserStore.UpdateUser(c.Context(), id, updateParams); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "user not found", Field: "error"})
		}
		return InternalServerError(c, ErrorObject{Msg: "failed to update user", Field: "error"})
	}

	return nil
}
