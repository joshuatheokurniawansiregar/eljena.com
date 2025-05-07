package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateCategory(ctx context.Context, categoryModel items.CategoriesModel) error {
	var query string = "INSERT INTO categories(category_name) VALUES($1)"
	var _, err = r.DB.ExecContext(ctx, query, categoryModel.CategoryName)
	if err != nil{
		return err
	}
	return nil
}