package items

type ItemTagRequest struct {
	ItemID int64 `json:"itemID"`
	TagID  int64 `json:"tagID"`
}

type ItemTagsModel struct {
	ItemID int64 `db:"item_id"`
	TagID  int64 `db:"tag_id"`
}