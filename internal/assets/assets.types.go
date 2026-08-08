package assets

type Mime string
type Ext string

type TypeResolver struct {
	Mime Mime
	Ext  Ext
}

const (
	//  Image Mime Allowed
	MimeImageJPEG Mime = "image/jpeg"
	MimeImagePNG  Mime = "image/png"
	MimeImageGIF  Mime = "image/gif"
	MimeImageWebP Mime = "image/webp"

	// Text Mime Allowed
	MimeTextCSV   Mime = "text/csv"
	MimeTextPlain Mime = "text/plain"
)
const (
	//  Image Mime Allowed
	ExtJpg  Ext = ".jpg"
	ExtPng  Ext = ".png"
	ExtGif  Ext = ".gif"
	ExtWebp Ext = ".webp"

	// Text Ext Allowed
	ExtCsv       Ext = ".csv"
	ExtTextPlain Ext = ".txt"
)

var AllowedImageType = map[string]TypeResolver{
	".jpg":  TypeResolver{Mime: MimeImageJPEG, Ext: ExtJpg},
	".png":  TypeResolver{Mime: MimeImagePNG, Ext: ExtPng},
	".gif":  TypeResolver{Mime: MimeImageGIF, Ext: ExtGif},
	".webp": TypeResolver{Mime: MimeImageWebP, Ext: ExtWebp},
}

var AllowedTextType = map[string]TypeResolver{
	".csv": TypeResolver{Mime: MimeTextCSV, Ext: ExtCsv},
	".txt": TypeResolver{Mime: MimeTextPlain, Ext: ExtTextPlain},
}

func isAllowedext(ext string) (TypeResolver, bool) {

	if mime, ok := AllowedImageType[ext]; ok {
		return mime, true
	}
	if mime, ok := AllowedTextType[ext]; ok {
		return mime, true
	}
	return TypeResolver{}, false
}

func isAllowedMimeByPolicy(payload AssetPolicy, mime Mime) bool {
	for _, m := range payload.AllowedMime {
		if m == mime {
			return true
		}
	}
	return false
}
