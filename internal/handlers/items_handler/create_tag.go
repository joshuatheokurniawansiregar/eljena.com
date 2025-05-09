package items_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (h *Handler) CreateTag(ctx *fiber.Ctx) error {
	var tagRequest items.TagRequest
	if err := ctx.BodyParser(&tagRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	if err := h.ItemsService.CreateTag(ctx.Context(), tagRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}
	
	return ctx.JSON(fiber.Map{
		"code": http.StatusCreated,
		"message": "Success  add a new tag",
		"data": tagRequest,
	})
}