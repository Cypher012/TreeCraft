package search

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func FindFolders(root, name string) ([]string, error) {
	var matches []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && strings.EqualFold(d.Name(), name) {
			matches = append(matches, path)
		}

		return nil
	})

	return matches, err
}
