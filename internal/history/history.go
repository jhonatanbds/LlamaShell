package history

import (
	"os"
	"os/user"
	"path/filepath"
)

func History() (string, error) {
	// Get the current user's home directory
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDir := usr.HomeDir

	// List of potential history files
	historyFiles := []string{
		filepath.Join(homeDir, ".bash_history"),
		filepath.Join(homeDir, ".zsh_history"),
		filepath.Join(homeDir, ".local/share/fish/fish_history"),
		filepath.Join(homeDir, ".config/fish/fish_history"),
		filepath.Join(homeDir, ".ash_history"),
		filepath.Join(homeDir, ".ksh_history"),
		filepath.Join(homeDir, ".tcsh_history"),
		filepath.Join(homeDir, ".history"), // generic fallback
	}

	historyFiles = filter(historyFiles, fileExists)

	content := ""
	// Check for the existence of each file and return the first one found
	for _, file := range historyFiles {
		if _, err := os.Stat(file); err == nil {
			data, err := os.ReadFile(file)
			if err != nil {
				return "", err
			}
			content += string(data)
		}
	}

	return content, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func filter(paths []string, f func(string) bool) []string {
	var filtered []string
	for _, path := range paths {
		if f(path) {
			filtered = append(filtered, path)
		}
	}
	return filtered
}
