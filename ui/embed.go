package ui

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var assetsDir embed.FS

// AssetsDirFS contains the embedded build directory files (without the "build" prefix)
var AssetsDirFS, _ = fs.Sub(assetsDir, "build")
