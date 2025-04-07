package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/my_gin_app/internal/model/items"
)



func (h *Handler) Service(ctx *fiber.Ctx) error {
	var categories []items.Category = items.GetItemsByCategoriesForServices()

	return ctx.Render("service", fiber.Map{
		"url": ctx.OriginalURL(),
		"categories": categories,
	})
}