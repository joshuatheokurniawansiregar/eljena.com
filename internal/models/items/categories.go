package items

type CategoriesModel struct {
	ID           int64  `db:"id"`
	CategoryName string `db:"category_name"`
}

type CategoriesRequest struct {
	CategoryName string `json:"categoryName"`
}

type CategoriesResponse struct {
	ID           int64  `json:"ID"`
	CategoryName string `json:"categoryName"`
}