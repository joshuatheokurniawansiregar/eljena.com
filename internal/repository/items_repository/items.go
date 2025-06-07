package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateItem(ctx context.Context, itemsModel items.ItemsModel) error {
	var query string = "INSERT INTO items(user_id, sub_category_id, item_name, item_description, item_image_url) VALUES($1, $2, $3, $4, $5);"
	var _, err = r.DB.ExecContext(ctx, query, itemsModel.UserID, itemsModel.SubCategoryID, itemsModel.ItemName, itemsModel.ItemDescription, itemsModel.ItemImageUrl)
	if err != nil{
		return err
	}
	return nil
}

func(r *Repository) GetItemsForCarousel(ctx context.Context)([]items.CarouselDataResponse, error){
	var query string = `
		SELECT i.id, i.item_name, i.item_description, i.item_image_url, s.sub_category_name, c.category_name, ARRAY_AGG(t.name) AS tags FROM items i
		JOIN sub_categories s ON i.sub_category_id = s.id
		JOIN categories c ON s.category_id = c.id
		JOIN item_tags it ON i.id = it.item_id
		JOIN tags t ON t.id = it.tag_id
		GROUP BY i.id, i.item_name, s.sub_category_name, c.category_name;
	`
	var carouselDataResponse []items.CarouselDataResponse = make([]items.CarouselDataResponse, 0)
	var rows, err = r.DB.QueryContext(ctx, query)
	if err != nil{
		return carouselDataResponse, err
	}
	defer rows.Close()
	for rows.Next(){
		var carouselDataModel items.CarouselDataModel
		var err error = rows.Scan(&carouselDataModel.ID, &carouselDataModel.ItemName, &carouselDataModel.ItemDescription,
		&carouselDataModel.ItemImageUrl, &carouselDataModel.SubCategoryName, &carouselDataModel.CategoryName, &carouselDataModel.Tags)

		if err != nil{
			return carouselDataResponse, err
		}

		carouselDataResponse = append(carouselDataResponse, items.CarouselDataResponse{
			ItemName: carouselDataModel.ItemName,
			ItemDescription: carouselDataModel.ItemDescription,
			ItemImageUrl: carouselDataModel.ItemImageUrl,
			SubCategoryName: carouselDataModel.SubCategoryName,
			CategoryName: carouselDataModel.CategoryName,
			Tags: carouselDataModel.Tags,
		})
	}
	return carouselDataResponse, nil
}