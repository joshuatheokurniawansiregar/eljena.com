package items

type ItemRequest struct {
	UserID        int64  `json:"userID"`
	SubCategoryID int64  `json:"subCategoryID"`
	ItemName      string `json:"itemName"`
}

type ItemsModel struct {
	ID            int64  `db:"id"`
	UserID        int64  `db:"user_id"`
	SubCategoryID int64  `db:"sub_category_id"`
	ItemName      string `db:"item_name"`
}

// type ItemsModel struct{

// }