package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Cypher012/TreeCraft/internal/fsutil"
	"github.com/Cypher012/TreeCraft/internal/input"
	"github.com/Cypher012/TreeCraft/internal/search"
)

func getWorkingDir() string {
	wd, _ := os.Getwd()
	return wd
}

func chooseFolder() string {
	query := input.ReadLine("Enter a folder name to search (press Enter to create path in root): ")
	if query == "" {
		return ""
	}

	if strings.ToLower(query) == "quit" || strings.ToLower(query) == "q" {
		os.Exit(0)
	}

	folders, _ := search.FindFolders(".", query)

	for i, folder := range folders {
		fmt.Printf("[%d] %s\n", i+1, folder)
	}

	choice := input.Get[int]("Choose a folder (default 1): ")
	return folders[choice-1]
}

func askPathParts() (string, string) {
	mode := input.Get[string]("Create path: ")
	clean := fsutil.CleanLeadingSlash(mode)
	return filepath.Split(clean)
}

func handleOverwrite(path string) {
	if fsutil.FileExists(path) {
		ok := input.Confirm("The file already exists. Overwrite?")
		if !ok {
			fmt.Println("Canceled")
			os.Exit(0)
		}
	}
}
