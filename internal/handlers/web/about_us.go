package web

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func Testt(context.Context){

}

func (h *Handler) AboutUs(ctx *fiber.Ctx)error {
	return ctx.Render("about_us", fiber.Map{
		"url": ctx.OriginalURL(),
	})
}