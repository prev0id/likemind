package website

import (
	"embed"
)

// type StaticFiles struct {
// 	fs embed.FS
// }

//go:embed static
var StaticFiles embed.FS
