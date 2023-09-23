package api

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed doc.swagger.json
var doc embed.FS

func OpenAPIHandler() http.Handler {
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(doc, ".")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}
