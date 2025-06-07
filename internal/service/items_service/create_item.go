package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateItem(ctx context.Context, itemRequest items.ItemRequest, itemImageUrl string) error {
	var itemsModel items.ItemsModel = items.ItemsModel{
		UserID: itemRequest.UserID,
		SubCategoryID: itemRequest.SubCategoryID,
		ItemName: itemRequest.ItemName,
		ItemDescription: itemRequest.ItemDescription,
		ItemImageUrl: itemImageUrl,
	}

	var err error = s.ItemsRepository.CreateItem(ctx, itemsModel)
	return err
}