package swaggerui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed html
var HtmlFs embed.FS

func HandleSwagger(router *mux.Router) {
	swaggerFs, err := fs.Sub(HtmlFs, "html")
	if err != nil {
		panic(err)
	}

	swaggerServer := http.FileServer(http.FS(swaggerFs))
	sh := http.StripPrefix("/swaggerui/", swaggerServer)
	router.PathPrefix("/swaggerui/").Handler(sh)
}
