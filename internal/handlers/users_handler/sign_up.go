package users_handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/users"
)

func (h *Handler) SignUp(ctx *fiber.Ctx) error {
	var context context.Context = ctx.Context()
	var request users.SignUpRequest
	if err := ctx.BodyParser(&request); err!=nil{
		return ctx.JSON(fiber.Map{
			"code":http.StatusBadRequest,
			"message":err.Error(),
			"data":nil,
		})
	}
	if err := h.UsersService.SignUp(context, request);err != nil{
		return ctx.JSON(fiber.Map{
			"code":http.StatusInternalServerError,
			"message":err.Error(),
			"data":nil,
		})
	}
	return ctx.JSON(fiber.Map{
		"code":http.StatusOK,
		"message":"User is created with email "+request.Email,
		"data":nil,
	})
}