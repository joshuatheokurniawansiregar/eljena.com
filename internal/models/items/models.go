package items

import "github.com/lib/pq"

type ItemRequest struct {
	UserID          int64  `form:"userID"`
	SubCategoryID   int64  `form:"subCategoryID"`
	ItemName        string `form:"itemName"`
	ItemDescription string `form:"itemDescription"`
}

type ItemsModel struct {
	ID              int64  `db:"id"`
	UserID          int64  `db:"user_id"`
	SubCategoryID   int64  `db:"sub_category_id"`
	ItemName        string `db:"item_name"`
	ItemDescription string `db:"item_description"`
	ItemImageUrl    string `db:"item_image_url"`
}

type CarouselDataResponse struct {
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	ItemImageUrl    string `json:"itemImageUrl"`
	SubCategoryName string `json:"subCategoryName"`
	CategoryName    string `json:"categoryName"`
	Tags            []string `json:"tags"`
}

type CarouselDataModel struct {
	ID              int64     `db:"id"`
	ItemName        string    `db:"item_name"`
	ItemDescription string    `db:"item_description"`
	ItemImageUrl    string    `db:"item_image_url"`
	SubCategoryName string    `db:"sub_category_name"`
	CategoryName    string    `db:"category_name"`
	Tags            pq.StringArray `db:"tags"`
}