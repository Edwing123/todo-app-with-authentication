package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed src/templates
	templates embed.FS

	//go:embed src/assets
	assets embed.FS
)

func createAssetsHandler() http.Handler {
	assetsFS, err := fs.Sub(assets, "src/assets")

	if err != nil {
		log.Fatalln(err)
	}

	assetsFileServerHandler := http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFS)))
	return assetsFileServerHandler
}
