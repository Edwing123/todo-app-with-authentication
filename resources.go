package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var (
	//go:embed src/templates
	templates embed.FS

	//go:embed src/assets
	assets embed.FS
)

func createAssetsHandler() http.Handler {
	serveAssetsFromDisk := flag.Bool(
		"serveAssetsFromDisk",
		false,
		"Use this flag to load assets from disk when developing",
	)

	flag.Parse()

	var assetsFS fs.FS
	var err error

	if *serveAssetsFromDisk {
		fmt.Println("ok")
		assetsFS = os.DirFS("src/assets")
	} else {
		assetsFS, err = fs.Sub(assets, "src/assets")
	}

	if err != nil {
		log.Fatalln(err)
	}

	assetsFileServerHandler := http.StripPrefix("/assets/", http.FileServer(http.FS(assetsFS)))
	return assetsFileServerHandler
}
