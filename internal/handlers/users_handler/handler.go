package users_handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/users"
)

type UsersService interface {
	SignUp(ctx context.Context, signUpRequest users.SignUpRequest) error
}

type Handler struct{
	*fiber.App
	UsersService UsersService
}

func NewHandler(app *fiber.App, usersService UsersService)*Handler{
	return &Handler{
		App: app,
		UsersService: usersService,
	}
}

func(h *Handler) RegisterRoute(){
	var routeGroup fiber.Router = h.Group("/api/v1")
	routeGroup.Post("/users", h.SignUp)
}