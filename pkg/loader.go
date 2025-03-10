package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// List of allowed Git hooks
var allowedHooks = map[string]bool{
	"applypatch-msg":     true,
	"commit-msg":         true,
	"post-update":        true,
	"pre-applypatch":     true,
	"pre-commit":         true,
	"pre-merge-commit":   true,
	"pre-push":           true,
	"pre-rebase":         true,
	"pre-receive":        true,
	"prepare-commit-msg": true,
	"push-to-checkout":   true,
	"update":             true,
}

// isValidHookType checks if the hookType is in the allowed list
func isValidHookType(hookType string) bool {
	return allowedHooks[hookType]
}

// LoadHook copies the hook script from .betterhook to .git/hooks
func LoadHook(hookType string) error {
	if !isValidHookType(hookType) {
		return fmt.Errorf("❌ Invalid hook type %q: only standard Git hooks are allowed", hookType)
	}

	// Check if .betterhook directory exists
	betterhookDir := ".betterhook"
	if _, err := os.Stat(betterhookDir); os.IsNotExist(err) {
		return fmt.Errorf("⚠️  .betterhook directory not found")
	}

	// Define script name and path
	scriptName := hookType + ".sh"
	srcPath := filepath.Join(betterhookDir, scriptName)

	// Check if the script exists in .betterhook
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("⚠️  Hook script %q not found in .betterhook folder", scriptName)
	}

	// Check if .git/hooks directory exists
	hooksDir := filepath.Join(".git", "hooks")
	if _, err := os.Stat(hooksDir); os.IsNotExist(err) {
		return fmt.Errorf("❌ '.git/hooks' directory not found; are you in a Git repository?")
	}

	// Define the destination path
	destPath := filepath.Join(hooksDir, hookType)

	// Copy the file content
	err := copyFile(srcPath, destPath)
	if err != nil {
		return fmt.Errorf("❌ Failed to copy hook: %w", err)
	}

	// Make the destination file executable
	if err := os.Chmod(destPath, 0755); err != nil {
		return fmt.Errorf("⚠️  Failed to set executable permission: %w", err)
	}

	fmt.Printf("✅ Hook %q successfully installed!\n", hookType)
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

// LoadAllHooks automatically installs all hook scripts from .betterhook directory
func LoadAllHooks() error {
	// Check if .betterhook directory exists
	betterhookDir := ".betterhook"
	if _, err := os.Stat(betterhookDir); os.IsNotExist(err) {
		return fmt.Errorf("⚠️  .betterhook directory not found")
	}

	// Check if .git/hooks directory exists
	hooksDir := filepath.Join(".git", "hooks")
	if _, err := os.Stat(hooksDir); os.IsNotExist(err) {
		return fmt.Errorf("❌ '.git/hooks' directory not found; are you in a Git repository?")
	}

	// Read all files in .betterhook directory
	files, err := os.ReadDir(betterhookDir)
	if err != nil {
		return fmt.Errorf("❌ Failed to read .betterhook directory: %w", err)
	}

	var errors []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Get hook type from filename (remove .sh extension)
		hookType := strings.TrimSuffix(file.Name(), ".sh")
		if !isValidHookType(hookType) {
			errors = append(errors, fmt.Sprintf("❌ Invalid hook type %q: only standard Git hooks are allowed", hookType))
			continue
		}

		// Define source and destination paths
		srcPath := filepath.Join(betterhookDir, file.Name())
		destPath := filepath.Join(hooksDir, hookType)

		// Copy the file content
		if err := copyFile(srcPath, destPath); err != nil {
			errors = append(errors, fmt.Sprintf("❌ Failed to copy hook %q: %v", hookType, err))
			continue
		}

		// Make the destination file executable
		if err := os.Chmod(destPath, 0755); err != nil {
			errors = append(errors, fmt.Sprintf("⚠️  Failed to set executable permission for %q: %v", hookType, err))
			continue
		}

		fmt.Printf("✅ Hook %q successfully installed!\n", hookType)
	}

	if len(errors) > 0 {
		return fmt.Errorf("Some hooks failed to install:\n%s", strings.Join(errors, "\n"))
	}

	return nil
}
