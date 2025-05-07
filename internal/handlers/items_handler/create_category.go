package items_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (h *Handler) CreateCategory(ctx *fiber.Ctx) error {
	var categoryRequest items.CategoriesRequest
	if err := ctx.BodyParser(&categoryRequest);err!=nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	if err := h.ItemsService.CreateCategory(ctx.Context(), categoryRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code": http.StatusInternalServerError,
		"message": "Category has been created",
		"data": nil,
	})
}