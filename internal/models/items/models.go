package items

type ItemRequest struct {
	UserID        int64  `form:"userID"`
	SubCategoryID int64  `form:"subCategoryID"`
	ItemName      string `form:"itemName"`
	// ItemDescription string `json:"itemDescription"`
}

type ItemsModel struct {
	ID            int64  `db:"id"`
	UserID        int64  `db:"user_id"`
	SubCategoryID int64  `db:"sub_category_id"`
	ItemName      string `db:"item_name"`
	// ItemDescription string `db:"item_description"`
	// ItemImageUrl    string `db:"item_image_url"`
}

// type ItemsModel struct{

// }