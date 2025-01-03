package static

import (
	_ "embed"

	"likemind/internal/domain"
)

var (
	//go:embed js/htmx.min.js
	scriptHTMX []byte
	//go:embed js/error_handler.js
	scriptErrorHandler []byte

	//go:embed css/styles.css
	styles []byte

	//go:embed img/favicon.png
	favicon []byte
	//go:embed img/img1.jpg
	testImg1 []byte
	//go:embed img/img2.jpg
	testImg2 []byte
)

var Files = map[string]FileDescriptor{
	domain.PathStaticHTMX: {
		File:     scriptHTMX,
		MimeType: mimeJS,
	},
	domain.PathStaticErrorHandler: {
		File:     scriptErrorHandler,
		MimeType: mimeJS,
	},
	domain.PathStaticStyles: {
		File:     styles,
		MimeType: mimeCSS,
	},
	domain.PathStaticFavicon: {
		File:     favicon,
		MimeType: mimePNG,
	},
	// TODO: remove
	"/static/test_image1.jpg": {
		File:     testImg1,
		MimeType: mimeJPEG,
	},
	"/static/test_image2.jpg": {
		File:     testImg2,
		MimeType: mimeJPEG,
	},
}

const (
	mimeJS   = "text/javascript"
	mimeCSS  = "text/css"
	mimeJPEG = "image/jpeg"
	mimePNG  = "image/png"
)

type FileDescriptor struct {
	File     []byte
	MimeType string
}
