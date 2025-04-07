package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joshuatheokurniawansiregar/my_gin_app/internal/handlers/api"
	"github.com/joshuatheokurniawansiregar/my_gin_app/internal/handlers/web"
)

func main() {
	// r := gin.Default()
	// r.Use(gin.Recovery())
	// r.Use(gin.Logger())
	// r.LoadHTMLGlob("views/*.tmpl")
	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "layout.tmpl", gin.H{
	// 		// "title":"hello world",
	// 		// "data":[]string{
	// 		// 	"joshua",
	// 		// 	"theo",
	// 		// 	"kurniawan",
	// 		// },
	// 		"url":ctx.Request.URL,
	// 	})
	// })
	// r.Run(":3000")

	engine:= html.New("./views", ".html")
	app:= fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// web route
	var handler *web.Handler = web.NewHandler(app)
	handler.RegisterRoute()

	// api route
	var apiHandler *api.Handler = api.NewHandler(app)
	apiHandler.RegisterRoute()

	app.Listen(":3000")
}