package app

import (
	"path/filepath"

	"github.com/Cypher012/TreeCraft/internal/fsutil"
)

func buildPath(wd, selected, dir, file string) string {
	if dir != "" {
		fullDir := filepath.Join(wd, selected, dir)
		fsutil.MakeDirs(fullDir)
	}

	return filepath.Join(wd, selected, dir, file)
}

func create(path string) {
	fsutil.CreateFile(path)
}
