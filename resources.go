package main

import "embed"

var (
	//go:embed src/templates
	templates embed.FS

	//go:embed src/assets
	assets embed.FS
)
