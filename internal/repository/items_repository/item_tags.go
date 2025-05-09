package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateItemTag(ctx context.Context, itemTagsModel items.ItemTagsModel) error {
	var query string = "INSERT INTO item_tags(item_id, tag_id) VALUES($1, $2)"
	var _, err = r.DB.ExecContext(ctx, query, itemTagsModel.ItemID, itemTagsModel.TagID)
	if err != nil{
		return err
	}

	return nil
}