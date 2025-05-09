package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateItemTag(context context.Context, itemTagRequest items.ItemTagRequest) error {
	var itemTagsModel items.ItemTagsModel = items.ItemTagsModel{
		ItemID: itemTagRequest.ItemID,
		TagID: itemTagRequest.TagID,
	}
	var err error = s.ItemsRepository.CreateItemTag(context, itemTagsModel)
	if err != nil{
		return err
	}
	return nil
}