package api

import (
	"errors"
	"net/http"

	"github.com/A3R0-01/head-hunter/db"
	"github.com/A3R0-01/head-hunter/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobPostHandler struct {
	store  *db.Store
	object string
}

func NewJobPostHandler(store *db.Store, object string) *JobPostHandler {
	return &JobPostHandler{
		store:  store,
		object: object,
	}
}

func (h *JobPostHandler) HandlePostJobPost(c *fiber.Ctx) error {
	var newJobPost types.CreateJobPostParams
	if err := c.BodyParser(&newJobPost); err != nil {
		return InternalServerError(c, ErrorObject{Msg: "invalid params", Field: "error"})
	}
	if errors := newJobPost.Validate(); len(errors) > 0 {
		return BadRequest(c, errors)
	}
	jobPost, err := newJobPost.FromParams()
	if err != nil {
		return InternalServerError(c, ErrorObject{Msg: "internal error", Field: "job post failed"})
	}
	jobPostCreated, err := h.store.JobPostStore.CreateJobPost(c.Context(), jobPost)
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(jobPostCreated)
}

func (h *JobPostHandler) HandleGetJobPost(c *fiber.Ctx) error {
	id := c.Params("id")

	jobPost, err := h.store.JobPostStore.GetJobPostByID(c.Context(), id)
	if err != nil {
		return BadRequest(c, ErrorObject{Msg: err.Error(), Field: "error"})
	}

	return c.JSON(jobPost)
}

func (h *JobPostHandler) HandleGetJobPosts(c *fiber.Ctx) error {
	jobPosts, err := h.store.GetJobPosts(c.Context(), db.Map{})
	if err != nil {
		return InternalServerError(c, err)
	}
	return c.JSON(jobPosts)
}

func (h *JobPostHandler) HandleDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.store.JobPostStore.DeleteJobPost(c.Context(), id); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "job post not found", Field: "error"})
		}
		return InternalServerError(c, err)
	}
	return c.Status(http.StatusOK).JSON("Job post deleted")
}

func (h *JobPostHandler) HandlePut(c *fiber.Ctx) error {
	var updateParams types.UpdateJobPostParams
	id := c.Params("id")
	if err := c.BodyParser(&updateParams); err != nil {
		return BadRequest(c, err)
	}

	if err := h.store.JobPostStore.UpdateJobPost(c.Context(), id, updateParams); err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return NotFound(c, ErrorObject{Msg: "job post not found", Field: "error"})
		}
		return InternalServerError(c, ErrorObject{Msg: "failed to update job post", Field: "error"})
	}

	return nil
}
