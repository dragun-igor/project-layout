package users

import (
	"context"
	"net/http"

	"github.com/dragun-igor/project-layout/internal/entities"
	"github.com/labstack/echo/v4"
)

type IService interface {
	Create(ctx context.Context, user entities.User) (entities.User, error)
}

type HTTPHandler struct {
	userService IService
}

func NewHTTPHandler(userService IService) *HTTPHandler {
	return &HTTPHandler{userService: userService}
}

func (h *HTTPHandler) Create(ctx echo.Context) error {
	var request struct {
		Login    string `json:"login" validate:"min=10"`
		Password string `json:"password" validate:"min=8"`
		Name     string `json:"required"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "binding error")
	}
	if err := ctx.Validate(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "validation error")
	}

	createdUser, err := h.userService.Create(ctx.Request().Context(), entities.User{
		Login:    request.Login,
		Password: request.Password,
		Name:     request.Name,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return ctx.JSON(http.StatusOK, createdUser)
}
