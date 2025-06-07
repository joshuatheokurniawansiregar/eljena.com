package items_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (s *Service) GetItemsForCarousel(ctx context.Context) ([]items.CarouselDataResponse, error) {
	var carouselDataResponse, err = s.ItemsRepository.GetItemsForCarousel(ctx)
	if err != nil{
		return carouselDataResponse, err
	}
	return carouselDataResponse, err
}