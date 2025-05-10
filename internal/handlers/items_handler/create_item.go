package items_handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (h *Handler) CreateItem(ctx *fiber.Ctx)error{
	var itemRequest items.ItemRequest
	
	if err := ctx.BodyParser(&itemRequest); err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	var file, err = ctx.FormFile("image")
	if err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	if err := h.ItemsService.CreateItem(ctx.Context(), itemRequest);err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	
	return ctx.JSON(fiber.Map{
		"code": http.StatusBadRequest,
		"message": "Create a new item succeed",
		"data": itemRequest,
	})
}