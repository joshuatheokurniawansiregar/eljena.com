package api

import "github.com/gofiber/fiber/v2"

type Handler struct {
	*fiber.App
}

func NewHandler(app *fiber.App) *Handler{
	return &Handler{
		App: app,
	}
}

func(h *Handler) RegisterRoute(){
	var routerGroup fiber.Router = h.Group("/api/v1")
	routerGroup.Get("/carousel", h.GetCarousel)
}