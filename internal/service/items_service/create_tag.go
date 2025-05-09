package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) CreateTag(context context.Context, tagRequest items.TagRequest) error{
	var tagsModel items.TagsModel = items.TagsModel{
		Name: tagRequest.Name,
	}
	var err error = s.ItemsRepository.CreateTag(context, tagsModel)
	if err != nil{
		return err
	}
	return nil
}