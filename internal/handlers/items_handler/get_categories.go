package items_handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetCategories(ctx *fiber.Ctx) error {
	var context context.Context = ctx.Context()

	var response, err = h.ItemsService.GetCategories(context)
	if err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data" : nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code": http.StatusOK,
		"message": "Data is shown",
		"data" : response,
	})
}