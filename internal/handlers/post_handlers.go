package handlers

import (
	"go_role/internal/core/domain"
	"go_role/internal/core/services"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler(ps *services.PostService) *PostHandler {
	return &PostHandler{postService: ps}
}

func (h *PostHandler) Create(c echo.Context) error {
	var create domain.PostCreate
	if err := c.Bind(&create); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	post, err := h.postService.CreatePost(c.Request().Context(), &create)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, post)
}
