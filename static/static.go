package static

import (
	"embed"
	"html/template"
	"os"
	"path/filepath"
)

//go:embed web/*
//go:embed web/_next/static
//go:embed web/_next/static/chunks/pages/*.js
//go:embed web/_next/static/*/*.js
var webFiles embed.FS

// GetWeb will return an embedded filesystem reference to the admin web app.
func GetWeb() embed.FS {
	return webFiles
}

//go:embed metadata.html.tmpl
var botMetadataTemplate embed.FS

// GetBotMetadataTemplate will return the bot/scraper metadata template.
func GetBotMetadataTemplate() (*template.Template, error) {
	name := "metadata.html.tmpl"
	t, err := template.ParseFS(botMetadataTemplate, name)
	tmpl := template.Must(t, err)
	return tmpl, err
}

//go:embed offline.ts
var offlineVideoSegment []byte

// GetOfflineSegment will return the offline video segment data.
func GetOfflineSegment() []byte {
	return getFileSystemStaticFileOrDefault("offline.ts", offlineVideoSegment)
}

func getFileSystemStaticFileOrDefault(path string, defaultData []byte) []byte {
	fullPath := filepath.Join("static", path)
	data, err := os.ReadFile(fullPath) //nolint: gosec
	if err != nil {
		return defaultData
	}

	return data
}
