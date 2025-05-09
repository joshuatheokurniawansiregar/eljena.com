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