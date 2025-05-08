package items_service

import (
	"context"
	"time"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateSubCategory(ctx context.Context, subCategoryRequest items.SubCategoriesRequest) error {
	var subCategoriesModel items.SubCategoriesModel = items.SubCategoriesModel{
		UserID: subCategoryRequest.UserID,
		CategoryID: subCategoryRequest.CategoryID,
		SubCategoryName: subCategoryRequest.SubCategoryName,
		CreatedAt: time.Now(),
	}

	return s.ItemsRepository.CreateSubCategory(ctx, subCategoriesModel)
}