package assets

type UploadPolicy struct {
	Category AssetCategory
	MaxSize  int64
	UserID   string
}

type AssetCategory string

const (
	CategoryProfile  AssetCategory = "profile"
	CategoryDocument AssetCategory = "document"
	CategoryPost     AssetCategory = "post"
	CategoryTemp     AssetCategory = "temp"
)

type AssetPolicy struct {
	Folder      string
	AllowedMime []Mime
	MaxSize     int64
}

var PolicyAsset = map[AssetCategory]AssetPolicy{
	CategoryProfile: {
		Folder:      "profile",
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG},
		MaxSize:     1024 * 1024 * 5,
	},
	CategoryDocument: {
		Folder:      "document",
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryPost: {
		Folder:      "post",
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryTemp: {
		Folder:      "temp",
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG},
		MaxSize:     1024 * 1024 * 1,
	},
}
