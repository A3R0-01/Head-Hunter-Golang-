package api

import (
	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
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

	return nil
}
