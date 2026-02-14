package fileops

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileMover_MoveFile(t *testing.T) {
	// Setup temporary directory
	tempDir, err := os.MkdirTemp("", "fileops-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir) // clean up

	// Create a dummy file
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(false, false)

	// Test moving to a new directory
	destDir := filepath.Join(tempDir, "Documents")
	newPath, err := mover.MoveFile(testFile, destDir)
	if err != nil {
		t.Fatalf("MoveFile failed: %v", err)
	}

	// Verify file exists in destination
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		t.Errorf("File was not moved to %s", newPath)
	}

	// Verify original file is gone
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Error("Original file still exists")
	}

	// Test duplicate handling
	// Create another file with same name
	testFile2 := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile2, []byte("newer content"), 0644); err != nil {
		t.Fatal(err)
	}

	// Try moving to same destination where "test.txt" already exists
	newPath2, err := mover.MoveFile(testFile2, destDir)
	if err != nil {
		t.Fatalf("MoveFile failed on duplicate: %v", err)
	}

	if newPath2 == newPath {
		t.Error("MoveFile should have renamed duplicate file, but returned same path")
	}

	// Verify both files exist
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		t.Error("First file disappeared")
	}
	if _, err := os.Stat(newPath2); os.IsNotExist(err) {
		t.Error("Renamed file not found")
	}
}

func TestFileMover_DryRun(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "dryrun-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "dryrun.txt")
	if err := os.WriteFile(testFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	mover := NewFileMover(true, false) // DryRun = true

	destDir := filepath.Join(tempDir, "DryRunFolder")
	_, err = mover.MoveFile(testFile, destDir)
	if err != nil {
		t.Fatalf("DryRun MoveFile failed: %v", err)
	}

	// Verify file was NOT moved
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Error("File should not have moved in DryRun mode")
	}

	// Verify destination directory was NOT created
	if _, err := os.Stat(destDir); !os.IsNotExist(err) {
		t.Error("Destination directory should not be created in DryRun mode")
	}
}
