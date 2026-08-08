package assets

import "time"

type AssetMetaData struct {
	ID        string      `json:"id"`
	Filename  string      `json:"filename"`
	FileKey   string      `json:"file_key"`
	Mime      Mime        `json:"mime"`
	Url       string      `json:"url"`
	Size      int64       `json:"size"`
	AuthorID  int         `json:"author_id"`
	Status    AssetStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
