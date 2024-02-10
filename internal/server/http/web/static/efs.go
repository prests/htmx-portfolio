package static

import (
	"embed"
	"fmt"

	"github.com/benbjohnson/hashfs"
)

//go:embed "js" "css" "img"
var assets embed.FS

//go:embed "docs"
var Docs embed.FS

type hashedFiles struct {
	files *hashfs.FS
}

var HashedFiles = &hashedFiles{hashfs.NewFS(assets)}

func (hf *hashedFiles) GetStaticFilePath(relativePath string) string {
	hashedPath := hf.files.HashName(relativePath)
	return fmt.Sprintf("static/%s", hashedPath)
}

func (hf *hashedFiles) GetFiles() *hashfs.FS {
	return hf.files
}
