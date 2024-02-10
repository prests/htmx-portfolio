package static

import (
	"embed"
	"fmt"

	"github.com/benbjohnson/hashfs"
)

//go:embed "js" "css" "img"
var files embed.FS

type hashedFiles struct {
	files *hashfs.FS
}

var HashedFiles = &hashedFiles{hashfs.NewFS(files)}

func (hf *hashedFiles) GetStaticFilePath(relativePath string) string {
	hashedPath := hf.files.HashName(relativePath)
	return fmt.Sprintf("static/%s", hashedPath)
}

func (hf *hashedFiles) GetFiles() *hashfs.FS {
	return hf.files
}
