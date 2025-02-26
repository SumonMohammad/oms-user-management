package sql

import (
	"embed"
	"io/fs"
)

//go:embed migrations
var migrations embed.FS

func GetMigrations() fs.FS {
	return migrations
}


