package website

import (
	"embed"
	"io"
	"path/filepath"
	"strings"

	"likemind/internal/domain/errs"
)

type StaticFiles struct {
	fs embed.FS
}

//go:embed static/*
var embeddedFiles embed.FS

func NewStaticFiles() *StaticFiles {
	return &StaticFiles{
		fs: embeddedFiles,
	}
}

type File struct {
	Reader io.ReadCloser
	MIME   string
}

func (s *StaticFiles) Get(path string) (*File, error) {
	const prefix = "static"
	if !strings.HasPrefix(path, prefix) {
		return nil, errs.NotFound("invalid static path: missing prefix")
	}

	f, err := s.fs.Open(path)
	if err != nil {
		return nil, errs.NotFound("file not found: " + path)
	}

	mimeType := mimeTypeByExtension(path)
	return &File{
		Reader: f,
		MIME:   mimeType,
	}, nil
}

func mimeTypeByExtension(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".jpeg", ".jpg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".ico":
		return "image/x-icon"
	default:
		return "application/octet-stream"
	}
}
