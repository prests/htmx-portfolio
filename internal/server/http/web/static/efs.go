package static

import "embed"

//go:embed "js" "css" "img"
var Files embed.FS
