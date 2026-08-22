package assets

import "time"

type AssetMetaData struct {
	ID        string        `json:"id"`
	ParentID  string        `json:"parent_id"`
	Filename  string        `json:"filename"`
	FileKey   string        `json:"file_key"`
	Mime      Mime          `json:"mime"`
	Size      int64         `json:"size"`
	AuthorID  string        `json:"author_id"`
	Status    AssetStatus   `json:"status"`
	Category  AssetCategory `json:"category"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}
