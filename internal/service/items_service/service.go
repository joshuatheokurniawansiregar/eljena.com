package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

type ItemsRepository interface {
	CreateCategory(ctx context.Context, categoryModel items.CategoriesModel)error
	CreateSubCategory(ctx context.Context, subCategoriesModel items.SubCategoriesModel) error
	CreateItem(ctx context.Context, itemsModel items.ItemsModel) error
	CreateTag(context context.Context, tagModel items.TagsModel) error
	CreateItemTag(ctx context.Context, itemTagsModel items.ItemTagsModel) error

	GetCategories(ctx context.Context) ([]items.CategoriesResponse , error)
	GetSubCategories(ctx context.Context)([]items.SubCategoriesResponse, error)
	GetItemsForCarousel()
}

type Service struct{
	ItemsRepository ItemsRepository
}

func NewService(itemsRepository ItemsRepository)*Service{
	return &Service{
		ItemsRepository: itemsRepository,
	}
}