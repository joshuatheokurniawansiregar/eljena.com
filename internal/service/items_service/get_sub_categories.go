package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) GetSubCategories(ctx context.Context) ([]items.SubCategoriesResponse, error) {
	var response, err = s.ItemsRepository.GetSubCategories(ctx)
	if err != nil{
		return response, err
	}

	return response, nil
}