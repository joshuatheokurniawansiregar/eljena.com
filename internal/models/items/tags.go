package items

type TagRequest struct {
	Name string `json:"name"`
}

type TagsModel struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
