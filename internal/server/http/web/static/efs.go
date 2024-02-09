package static

import (
	"embed"

	"github.com/benbjohnson/hashfs"
)

//go:embed "js" "css" "img"
var files embed.FS

var HashedFiles = hashfs.NewFS(files)
