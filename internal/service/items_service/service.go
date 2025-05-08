package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

type ItemsRepository interface {
	CreateCategory(ctx context.Context, categoryModel items.CategoriesModel)error
	CreateSubCategory(ctx context.Context, subCategoriesModel items.SubCategoriesModel) error
}

type Service struct{
	ItemsRepository ItemsRepository
}

func NewService(itemsRepository ItemsRepository)*Service{
	return &Service{
		ItemsRepository: itemsRepository,
	}
}