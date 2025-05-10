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

func(r *Repository) GetSubCategories(ctx context.Context)([]items.SubCategoriesResponse, error){
	var subCategoriesResponse []items.SubCategoriesResponse = make([]items.SubCategoriesResponse, 0)
	var query string = "SELECT * FROM sub_categories;"
	var rows, err = r.DB.QueryContext(ctx, query)
	if err != nil{
		return subCategoriesResponse, err
	}

	defer rows.Close()

	for rows.Next(){
		var subCategoriesModel items.SubCategoriesModel
		var err error = rows.Scan(&subCategoriesModel.ID, &subCategoriesModel.UserID, &subCategoriesModel.CategoryID, &subCategoriesModel.SubCategoryName, &subCategoriesModel.CreatedAt, &subCategoriesModel.UpdatedAt)

		if err != nil{
			return subCategoriesResponse, err
		}

		var response items.SubCategoriesResponse = items.SubCategoriesResponse{
			ID: subCategoriesModel.ID,
			UserID: subCategoriesModel.UserID,
			CategoryID: subCategoriesModel.CategoryID,
			SubCategoryName: subCategoriesModel.SubCategoryName,
			CreatedAt: subCategoriesModel.CreatedAt,
			UpdatedAt: subCategoriesModel.UpdatedAt,
		}

		subCategoriesResponse = append(subCategoriesResponse, response)
	}

	return subCategoriesResponse, nil
}