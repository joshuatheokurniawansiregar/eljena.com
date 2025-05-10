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

	GetCategories(ctx context.Context) ([]items.CategoriesResponse, error)
	GetSubCategories(ctx context.Context) ([]items.SubCategoriesResponse, error)
	GetItemsForCarousel()
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
	groupRouter.Get("/categories", h.GetCategories)

	groupRouter.Post("/sub_categories", h.CreateSubCategory)
	groupRouter.Get("/sub_categories", h.GetSubCategories)

	groupRouter.Post("/items", h.CreateItem)
	groupRouter.Post("/tags", h.CreateTag)
	groupRouter.Post("/item_tags", h.CreateItemTag)

}