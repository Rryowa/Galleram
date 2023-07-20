package templates

import "embed"

// FS embedding files in binary
//
//go:embed *.gohtml
var FS embed.FS
