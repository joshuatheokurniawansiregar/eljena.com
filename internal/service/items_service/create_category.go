package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateCategory(ctx context.Context, categoryReq items.CategoriesRequest)error {
	var categoryModel items.CategoriesModel = items.CategoriesModel{
		ID: 0,
		CategoryName: categoryReq.CategoryName,
	}

	var err error = s.ItemsRepository.CreateCategory(ctx, categoryModel)
	return err
}