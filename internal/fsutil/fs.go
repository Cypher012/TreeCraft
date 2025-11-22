package fsutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CleanLeadingSlash(s string) string {
	return strings.TrimPrefix(s, "/")
}

func MakeDirs(path string) error {
	return os.MkdirAll(path, 0755)
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func OpenInEditor(path string) {
	editor := "cursor"

	// Fallback to VS Code if Cursor isn't installed
	if _, err := exec.LookPath(editor); err != nil {
		editor = "code"
	}

	cmd := exec.Command(editor, path)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Could not open editor:", err)
	}
}

func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	f.Close()

	OpenInEditor(path)
}
