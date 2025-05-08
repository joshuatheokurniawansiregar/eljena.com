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
	UserID          int64  `db:"user_id"`
	CategoryID      int64  `db:"category_id"`
	SubCategoryName string `db:"sub_category_name"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}