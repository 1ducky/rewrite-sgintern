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
		Folder:      string(CategoryProfile),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 5,
	},
	CategoryDocument: {
		Folder:      string(CategoryDocument),
		AllowedMime: []Mime{MimeDocPDF, MimeDocDocx, MimeDocXlsx, MimeDocPptx, MimeDocDoc},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryPost: {
		Folder:      string(CategoryPost),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 10,
	},
	CategoryTemp: {
		Folder:      string(CategoryTemp),
		AllowedMime: []Mime{MimeImageJPEG, MimeImagePNG, MimeImageGIF, MimeImageWebP},
		MaxSize:     1024 * 1024 * 1,
	},
}
