package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/my_gin_app/internal/model/items"
)



func (h *Handler) Program(ctx *fiber.Ctx) error {
	var categories []items.Category = items.GetItemsByCategoriesForProgram()

	return ctx.Render("program", fiber.Map{
		"url": ctx.OriginalURL(),
		"categories": categories,
	})
}