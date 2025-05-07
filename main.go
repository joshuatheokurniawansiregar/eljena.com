package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/joshuatheokurniawansiregar/eljena/internal/handlers/api"
	"github.com/joshuatheokurniawansiregar/eljena/internal/handlers/items_handler"
	"github.com/joshuatheokurniawansiregar/eljena/internal/handlers/users_handler"
	"github.com/joshuatheokurniawansiregar/eljena/internal/handlers/web"
	"github.com/joshuatheokurniawansiregar/eljena/internal/repository/items_repository"
	"github.com/joshuatheokurniawansiregar/eljena/internal/repository/users_repository"
	"github.com/joshuatheokurniawansiregar/eljena/internal/service/items_service"
	"github.com/joshuatheokurniawansiregar/eljena/internal/service/users_service"
	"github.com/joshuatheokurniawansiregar/eljena/pkg/internalsql"
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
	godotenv.Load(".env")
	var dataSourceName string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_PASSSWORD"),os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT"), os.Getenv("POSTGRESQL_DB"))

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

	var db, err = internalsql.Connect(dataSourceName)
	if err != nil{
		log.Println(err.Error())
	}

	defer db.Close()

	var usersRepository *users_repository.Repository = users_repository.NewRepository(db)
	var usersService *users_service.Service = users_service.NewService(usersRepository)
	var usersHandler *users_handler.Handler = users_handler.NewHandler(app, usersService)
	usersHandler.RegisterRoute()

	var itemsRepository *items_repository.Repository = items_repository.NewRepository(db)
	var itemsService *items_service.Service = items_service.NewService(itemsRepository)
	var itemsHandler *items_handler.Handler = items_handler.NewHandler(app, itemsService)
	itemsHandler.RegisterRoute()

	app.Listen(":3000")
}