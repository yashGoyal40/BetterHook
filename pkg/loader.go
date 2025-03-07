package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func LoadHook(hookType string) error {
	if hookType != "pre-commit" && hookType != "pre-push" {
		return fmt.Errorf("invalid hook type %q: only 'pre-commit' and 'pre-push' are allowed", hookType)
	}

	// Check if .betterhook directory exists
	betterhookDir := ".betterhook"
	if _, err := os.Stat(betterhookDir); os.IsNotExist(err) {
		return fmt.Errorf(".betterhook directory not found")
	}

	// Define script name and path
	scriptName := hookType + ".sh"
	srcPath := filepath.Join(betterhookDir, scriptName)

	// Check if the script exists in .betterhook
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("hook script %q not found in .betterhook folder", scriptName)
	}

	// Check if .git/hooks directory exists
	hooksDir := filepath.Join(".git", "hooks")
	if _, err := os.Stat(hooksDir); os.IsNotExist(err) {
		return fmt.Errorf("'.git/hooks' directory not found; are you in a git repository?")
	}

	// Define the destination path
	destPath := "./.git/hooks/" + hookType

	// Copy the file content
	err := copyFile(srcPath, destPath)
	if err != nil {
		return fmt.Errorf("failed to copy hook: %w", err)
	}

	// Make the destination file executable
	if err := os.Chmod(destPath, 0755); err != nil {
		return fmt.Errorf("failed to set executable permission: %w", err)
	}

	return nil
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}
