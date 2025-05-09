package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateTag(context context.Context, tagModel items.TagsModel)error {
	var query string = "INSERT INTO tags(name) VALUES($1)"
	var _, err = r.DB.ExecContext(context, query, tagModel.Name)
	if err != nil{
		return err
	}
	return nil
}

