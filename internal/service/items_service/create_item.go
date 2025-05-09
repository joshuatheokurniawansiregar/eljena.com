package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateItem(ctx context.Context, itemRequest items.ItemRequest) error {
	var itemsModel items.ItemsModel = items.ItemsModel{
		UserID: itemRequest.UserID,
		SubCategoryID: itemRequest.SubCategoryID,
		ItemName: itemRequest.ItemName,
	}

	var err error = s.ItemsRepository.CreateItem(ctx, itemsModel)
	return err
}