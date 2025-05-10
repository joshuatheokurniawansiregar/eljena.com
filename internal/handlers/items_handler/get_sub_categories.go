package items_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetSubCategories(ctx *fiber.Ctx)error{
	var subCategoriesResponse, err = h.ItemsService.GetSubCategories(ctx.Context())
	if err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message":err.Error(),
			"data":nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code": http.StatusOK,
		"message": "Data is shown",
		"data":subCategoriesResponse,
	})
}