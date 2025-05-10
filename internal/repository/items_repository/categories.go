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

func(r *Repository) GetCategories(ctx context.Context) ([]items.CategoriesResponse , error){
	var categoriesResponse []items.CategoriesResponse = make([]items.CategoriesResponse, 0)
	var query string = "SELECT * FROM categories;"
	var rows, err = r.DB.QueryContext(ctx, query)
	if err != nil{
		return categoriesResponse, err
	}

	defer rows.Close()

	for rows.Next(){
		var categoriesModel items.CategoriesModel
		var err error = rows.Scan(&categoriesModel.ID, &categoriesModel.CategoryName)
		if err != nil{
			return categoriesResponse, err
		}

		var response = items.CategoriesResponse{
			ID: categoriesModel.ID,
			CategoryName: categoriesModel.CategoryName,
		}

		categoriesResponse = append(categoriesResponse, response)
	}

	return categoriesResponse, nil
}