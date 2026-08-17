package assets

type RecordPayload struct {
	FileKey  string
	Status   AssetStatus
	Filename string
	Mime     Mime
	Path     string
	AuthorID string
	Category AssetCategory
}

type UpdatePayload struct {
	Id        string
	Size      int64
	Status    AssetStatus
	OldStatus AssetStatus
	AuthorID  string
}

type DeletePayload struct {
	IDs []string
}

type DeleteReaport struct {
	id  string
	err error
}
