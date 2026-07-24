package assets

type Mime string

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

var AllowedImageMime = map[Mime]string{
	MimeImageJPEG: ".jpg",
	MimeImagePNG:  ".png",
	MimeImageGIF:  ".gif",
	MimeImageWebP: ".webp",
}

var AllowedTextMime = map[Mime]string{
	MimeTextCSV:   ".csv",
	MimeTextPlain: ".txt",
}
