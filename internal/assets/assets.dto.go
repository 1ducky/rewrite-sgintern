package assets

type RecordPayload struct {
	Status   AssetStatus
	Filename string
	Mime     Mime
	Url      string
	AuthorID string
}

type UpdatePayload struct {
	Id        string
	Size      int64
	Status    AssetStatus
	OldStatus AssetStatus
	AuthorID  string
}
