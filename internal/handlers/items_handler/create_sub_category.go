package items_handler

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (h *Handler) CreateSubCategory(ctx *fiber.Ctx) error {
	var context context.Context = ctx.Context()
	var subCategoryRequest items.SubCategoriesRequest
	if err := ctx.BodyParser(&subCategoryRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}
	var err error = h.ItemsService.CreateSubCategory(context, subCategoryRequest)
	if err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"code": http.StatusCreated,
		"message": "Create a new sub category succeed",
		"data": subCategoryRequest,
	})
}