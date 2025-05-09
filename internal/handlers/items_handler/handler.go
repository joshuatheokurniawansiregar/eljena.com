package items_handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

type ItemsService interface {
	CreateCategory(ctx context.Context, categoryReq items.CategoriesRequest)error
	CreateSubCategory(ctx context.Context, subCategoryRequest items.SubCategoriesRequest) error
	CreateItem(ctx context.Context, itemRequest items.ItemRequest) error
	CreateTag(context context.Context, tagRequest items.TagRequest) error
	CreateItemTag(context context.Context, itemTagRequest items.ItemTagRequest) error
}

type Handler struct{
	*fiber.App
	ItemsService ItemsService
}

func NewHandler(app *fiber.App, itemsService ItemsService)*Handler{
	return &Handler{
		App: app,
		ItemsService: itemsService,
	}
}

func(h *Handler) RegisterRoute(){
	var groupRouter fiber.Router = h.Group("/api/v1")
	groupRouter.Post("/categories", h.CreateCategory)
	groupRouter.Post("/sub_categories", h.CreateSubCategory)
	groupRouter.Post("/items", h.CreateItem)
	groupRouter.Post("/tags", h.CreateTag)
	groupRouter.Post("/item_tags", h.CreateItemTag)
}