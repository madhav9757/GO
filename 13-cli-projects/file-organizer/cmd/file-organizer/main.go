package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"file-organizer/internal/app"
	"file-organizer/internal/classifier"
	"file-organizer/internal/fileops"
)

func main() {
	// 1. Definition of flags
	dirPtr := flag.String("dir", ".", "Directory to organize")
	dryRunPtr := flag.Bool("dry-run", false, "Preview mode: Show actions without executing")
	verbosePtr := flag.Bool("v", false, "Enable verbose logging")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: file-organizer [options] [directory]\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// 2. Resolve directory
	targetDir := *dirPtr
	if flag.NArg() > 0 {
		targetDir = flag.Arg(0)
	}

	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatalf("❌ Error resolving path: %v", err)
	}

	// 3. Initialize dependencies
	clf := classifier.NewClassifier()
	mover := fileops.NewFileMover(*dryRunPtr, *verbosePtr)

	// 4. Configure Application
	application := &app.App{
		Config: app.Config{
			SourceDir: absPath,
			DryRun:    *dryRunPtr,
			Verbose:   *verbosePtr,
		},
		Classifier: clf,
		Mover:      mover,
	}

	// 5. Execution
	if err := application.Run(); err != nil {
		log.Fatalf("❌ Execution failed: %v", err)
	}
}
