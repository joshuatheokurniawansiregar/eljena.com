package items_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (h *Handler) CreateItemTag(ctx *fiber.Ctx) error{
	var itemTagRequest items.ItemTagRequest

	if err := ctx.BodyParser(&itemTagRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	if err := h.ItemsService.CreateItemTag(ctx.Context(), itemTagRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code": http.StatusCreated,
		"message": "Success add a new item tag",
		"data": itemTagRequest,
	})
}