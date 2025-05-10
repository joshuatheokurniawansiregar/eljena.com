package items

import (
	"database/sql"
	"time"
)

type SubCategoriesRequest struct {
	UserID          int64  `json:"userID"`
	CategoryID      int64  `json:"categoryID"`
	SubCategoryName string `json:"subCategoryName"`
}

type SubCategoriesModel struct {
	ID 				int64 `db:"id"`
	UserID          int64  `db:"user_id"`
	CategoryID      int64  `db:"category_id"`
	SubCategoryName string `db:"sub_category_name"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt 		sql.NullTime `db:"updated_at"`
}

type SubCategoriesResponse struct {
	ID           	int64  			`json:"ID"`
	UserID 	     	int64 			`json:"userID"`
	CategoryID   	int64 			`json:"categoryID"`
	SubCategoryName string 			`json:"subCategoryName"`
	CreatedAt 		time.Time 		`json:"createdAt"`
	UpdatedAt		sql.NullTime	`json:"updatedAt"`
	CategoryName 	string 			`json:"categoryName"`
}