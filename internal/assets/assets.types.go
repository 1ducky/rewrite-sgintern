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

var AllowedImageType = map[Mime]TypeResolver{
	MimeImageJPEG: {Mime: MimeImageJPEG, Ext: ExtJpg},
	MimeImagePNG:  {Mime: MimeImagePNG, Ext: ExtPng},
	MimeImageGIF:  {Mime: MimeImageGIF, Ext: ExtGif},
	MimeImageWebP: {Mime: MimeImageWebP, Ext: ExtWebp},
}

var AllowedTextType = map[Mime]TypeResolver{
	MimeTextCSV:   {Mime: MimeTextCSV, Ext: ExtCsv},
	MimeTextPlain: {Mime: MimeTextPlain, Ext: ExtTextPlain},
}

func isAllowedExt(mime Mime) (TypeResolver, bool) {

	if mime, ok := AllowedImageType[mime]; ok {
		return mime, true
	}
	if mime, ok := AllowedTextType[mime]; ok {
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
