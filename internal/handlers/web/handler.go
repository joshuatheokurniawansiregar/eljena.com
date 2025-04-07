package web

import "github.com/gofiber/fiber/v2"

type Handler struct {
	*fiber.App
}

func NewHandler(app *fiber.App)*Handler{
	return &Handler{
		App: app,
	}
}

func(h *Handler) RegisterRoute(){
	h.Static("assets", "./public")
	h.Get("/", h.Home)
	h.Get("/program", h.Program)
	h.Get("/service", h.Service)
	h.Get("/about_us", h.AboutUs)
}