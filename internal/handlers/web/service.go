package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)



func (h *Handler) Service(ctx *fiber.Ctx) error {
	var categories []items.Category = items.GetItemsByCategoriesForServices()

	return ctx.Render("service", fiber.Map{
		"url": ctx.OriginalURL(),
		"categories": categories,
	})
}