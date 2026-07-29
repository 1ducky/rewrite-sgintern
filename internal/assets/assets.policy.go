package assets

type UploadPolicy struct {
	AllowdFileExt   []string
	AllowedMimeType []Mime
	MaxSize         int64
	Directory       string
	FileNameLength  int
}
