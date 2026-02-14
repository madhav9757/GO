package fileops

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileMover handles moving files with safety checks and dry-run capability
type FileMover struct {
	DryRun  bool
	Verbose bool
}

// NewFileMover creates a new FileMover instance
func NewFileMover(dryRun, verbose bool) *FileMover {
	return &FileMover{
		DryRun:  dryRun,
		Verbose: verbose,
	}
}

// MoveFile moves a file from src to a directory defined by destDir.
// It handles duplicate naming automatically.
func (m *FileMover) MoveFile(srcPath, destDir string) (string, error) {
	fileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, fileName)

	// Ensure destination directory exists (unless dry run, but even then good to simulate checks if possible,
	// but failing on MkdirAll is real error).
	if !m.DryRun {
		if err := os.MkdirAll(destDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", destDir, err)
		}
	}

	// Check for duplicates
	if _, err := os.Stat(destPath); err == nil {
		// File exists, generate unique name
		ext := filepath.Ext(fileName)
		nameWithoutExt := strings.TrimSuffix(fileName, ext)
		timestamp := time.Now().UnixNano()
		newFileName := fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)
		destPath = filepath.Join(destDir, newFileName)

		if m.Verbose {
			log.Printf("⚠️  Duplicate found: %s -> Renaming to %s", fileName, newFileName)
		}
	}

	// Perform the move
	if m.Verbose || m.DryRun {
		action := "Moving"
		if m.DryRun {
			action = "[DRY RAP] Would move"
		}
		log.Printf("👉 %s: %s -> %s", action, filepath.Base(srcPath), destPath)
	}

	if !m.DryRun {
		if err := os.Rename(srcPath, destPath); err != nil {
			// Cross-device link error handling (optional, but good for production)
			// standard os.Rename might fail across partitions. copy+delete is fallback.
			// For now assuming same filesystem, but let's add simple copy-delete fallback if needed.
			// kept simple for now as per typical use case.
			return "", fmt.Errorf("failed to move file: %w", err)
		}
	}

	return destPath, nil
}

// CopyFile is a helper if we implement cross-device move fallback later
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
