package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func(s *Service) GetCategories(ctx context.Context) ([]items.CategoriesResponse, error) {
	var response, err = s.ItemsRepository.GetCategories(ctx)
	if err != nil{
		return response, err
	}

	return response, nil
}