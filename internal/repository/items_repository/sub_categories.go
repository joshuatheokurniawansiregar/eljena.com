package items_repository

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func (r *Repository) CreateSubCategory(ctx context.Context, subCategoriesModel items.SubCategoriesModel) error {
	var query string = "INSERT INTO sub_categories(user_id, category_id, sub_category_name, created_at) VALUES($1, $2, $3, $4);"
	var _, err = r.DB.ExecContext(ctx, query, subCategoriesModel.UserID, subCategoriesModel.CategoryID, subCategoriesModel.SubCategoryName, subCategoriesModel.CreatedAt)
	if err != nil{
		return err
	}
	return nil
}