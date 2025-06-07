package items_handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetItemsForCarousel(ctx *fiber.Ctx) error {
	var carouselDataResponse, err = h.ItemsService.GetItemsForCarousel(ctx.Context())
	if err != nil{
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"code": http.StatusOK,
		"message": "Show all items for carousel",
		"data": carouselDataResponse,
	})
}