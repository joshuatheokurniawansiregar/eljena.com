package web

import "github.com/gofiber/fiber/v2"

func (h *Handler) AboutUs(ctx *fiber.Ctx)error {

	return ctx.Render("about_us", fiber.Map{
		"url": ctx.OriginalURL(),
	})
}