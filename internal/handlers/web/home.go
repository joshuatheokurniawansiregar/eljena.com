package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)



func (h *Handler) Home(ctx *fiber.Ctx) error {
	var categories []items.Category = items.GetItemsByCategories()

	return ctx.Render("home", fiber.Map{
		"url": ctx.OriginalURL(),
		"categories": categories,
	})
}