package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateItem(ctx context.Context, itemsModel items.ItemsModel) error {
	var query string = "INSERT INTO items(user_id, sub_category_id, item_name) VALUES($1, $2, $3);"
	var _, err = r.DB.ExecContext(ctx, query, itemsModel.UserID, itemsModel.SubCategoryID, itemsModel.ItemName)
	if err != nil{
		return err
	}
	return nil
}

func(r *Repository) GetItemsForCarousel(){
	// var query string = `
	// 	SELECT i.id, i.item_name, s.sub_category_name, c.category_name, ARRAY_AGG(t.name) AS tags FROM items i
	// 	FULL JOIN sub_categories s ON i.sub_category_id = s.id
	// 	FULL JOIN categories c ON s.category_id = c.id
	// 	FULL JOIN item_tags it ON i.id = it.item_id
	// 	FULL JOIN tags t ON t.id = it.tag_id
	// 	GROUP BY i.id, i.item_name, s.sub_category_name, c.category_name;
	// `
}