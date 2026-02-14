package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"file-organizer/internal/classifier"
	"file-organizer/internal/fileops"
)

// App holds the application state
type App struct {
	Config     Config
	Classifier *classifier.Classifier
	Mover      *fileops.FileMover
}

// Config for the application
type Config struct {
	SourceDir string
	DryRun    bool
	Verbose   bool
}

// Stats tracks execution metrics
type Stats struct {
	FilesMoved   int64
	FilesSkipped int64
	FilesFailed  int64
	TotalSize    int64 // bytes
	Duration     time.Duration
}

// Run executes the file organization process
func (a *App) Run() error {
	start := time.Now()

	// Validate source directory
	info, err := os.Stat(a.Config.SourceDir)
	if err != nil {
		return fmt.Errorf("failed to access source directory: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("source path is not a directory: %s", a.Config.SourceDir)
	}

	fmt.Printf("🚀 Starting file organization in: %s\n", a.Config.SourceDir)
	if a.Config.DryRun {
		fmt.Println("🚧 DRY RUN MODE: No changes will be made.")
	}

	// Prepare channels for worker pool
	type fileTask struct {
		path string
		info os.DirEntry
	}

	filesCh := make(chan fileTask, 100)
	var wg sync.WaitGroup
	var stats Stats

	// Start workers
	numWorkers := runtime.NumCPU() * 2 // Parallelism for faster processing
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range filesCh {
				a.processFile(task.path, task.info, &stats)
			}
		}()
	}

	// Walk directory and feed workers
	fmt.Println("Scanning files...")
	entries, err := os.ReadDir(a.Config.SourceDir)
	if err != nil {
		close(filesCh)
		return fmt.Errorf("failed to read directory contents: %w", err)
	}

	for _, entry := range entries {
		// Skip directories and hidden files
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		// Skip the executable itself
		exePath, err := os.Executable()
		if err == nil {
			if filepath.Base(exePath) == entry.Name() {
				continue
			}
		}

		// Send to channel
		filesCh <- fileTask{
			path: filepath.Join(a.Config.SourceDir, entry.Name()),
			info: entry,
		}
	}
	close(filesCh) // Signal workers to finish

	wg.Wait()
	stats.Duration = time.Since(start)

	// Print Summary
	a.printSummary(stats)
	return nil
}

func (a *App) processFile(path string, entry os.DirEntry, stats *Stats) {
	// Re-check classification
	cat := a.Classifier.Classify(entry.Name())
	destDir := filepath.Join(a.Config.SourceDir, string(cat))

	// Move file
	if _, err := a.Mover.MoveFile(path, destDir); err != nil {
		if a.Config.Verbose {
			log.Printf("❌ Failed to move %s: %v", entry.Name(), err)
		}
		atomic.AddInt64(&stats.FilesFailed, 1)
	} else {
		atomic.AddInt64(&stats.FilesMoved, 1)
		// Safely get size
		// Safely get size
		if info, err := entry.Info(); err == nil {
			atomic.AddInt64(&stats.TotalSize, info.Size())
		}
	}
}

func (a *App) printSummary(s Stats) {
	fmt.Println("\n📊 Execution Summary")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("⏱️  Duration:      %v\n", s.Duration)
	fmt.Printf("✅ Files Moved:   %d\n", atomic.LoadInt64(&s.FilesMoved))
	fmt.Printf("⏭️  Skipped/Immu:  %d (Implied if total - moved)\n", atomic.LoadInt64(&s.FilesSkipped)) // Need to track skipped logic properly if I want this metric
	fmt.Printf("❌ Failed:        %d\n", atomic.LoadInt64(&s.FilesFailed))
	fmt.Printf("💾 Total Size:    %.2f MB\n", float64(atomic.LoadInt64(&s.TotalSize))/(1024*1024))
	fmt.Println("--------------------------------------------------")
}
