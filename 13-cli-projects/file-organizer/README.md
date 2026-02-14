# 📂 Go File Organizer: The Ultimate Guide

> **"Order is the sanity of the mind, the health of the body, the peace of the city, the security of the state. As the beams to a house, as the bones to the microcosm of man, so is order to all things."**  
> — _Robert Southey_

---

## 📖 Table of Contents

1. [Introduction](#-introduction)
2. [Project Philosophy](#-project-philosophy)
3. [Features at a Glance](#-features-at-a-glance)
4. [Architecture & Design](#-architecture--design)
   - [Directory Structure](#directory-structure)
   - [Component Breakdown](#component-breakdown)
   - [Data Flow](#data-flow)
5. [Deep Dive: Implementation Details](#-deep-dive-implementation-details)
   - [The Entry Point (`cmd/file-organizer`)](#the-entry-point-cmdfile-organizer)
   - [Application Logic (`internal/app`)](#application-logic-internalapp)
   - [The Classifier Engine (`internal/classifier`)](#the-classifier-engine-internalclassifier)
   - [Safe File Operations (`internal/fileops`)](#safe-file-operations-internalfileops)
6. [Concurrency Model](#-concurrency-model)
   - [Worker Pools](#worker-pools)
   - [Synchronization](#synchronization)
7. [Safety Mechanisms](#-safety-mechanisms)
   - [Dry Run Mode](#dry-run-mode)
   - [Duplicate Collision Handling](#duplicate-collision-handling)
   - [Atomic Stats](#atomic-stats)
8. [Installation Guide](#-installation-guide)
   - [Prerequisites](#prerequisites)
   - [Building from Source](#building-from-source)
9. [Usage Manual](#-usage-manual)
   - [Command Line Flags](#command-line-flags)
   - [Examples](#examples)
10. [Supported File Types](#-supported-file-types)
11. [Testing & Quality Assurance](#-testing--quality-assurance)
12. [Troubleshooting](#-troubleshooting)
13. [Future Roadmap](#-future-roadmap)
14. [Contributing](#-contributing)
15. [License](#-license)

---

## 🌟 Introduction

Welcome to **Go File Organizer**, a robust, high-performance command-line interface (CLI) tool designed to bring order to the chaos of your digital filesystem.

In the modern digital age, our "Downloads" and "Documents" folders often become dumping grounds for thousands of files—images, PDFs, installers, scripts, and archives—intermingled in a messy heap. Finding what you need becomes a chore.

**Go File Organizer** solves this problem by automatically scanning a target directory, identifying files by their extensions, and moving them into categorized subfolders (e.g., `Images/`, `Documents/`, `Code/`). It does this safely, swiftly, and intelligently.

Built with **Go (Golang)**, this tool leverages the language's strong static typing, compilation to native binaries, and legendary concurrency primitives to process thousands of files in mere seconds.

---

## 🧠 Project Philosophy

This project is built on three core pillars:

1.  **Safety First**: A file organizer should _never_ lose data. We implement strict "Dry Run" capabilities, duplicate detection (renaming instead of overwriting), and careful error handling.
2.  **Performance**: We don't just move files one by one. We utilize Go's Goroutines to parallelize file processing, ensuring the tool scales to handle massive directories with ease.
3.  **Maintainability**: The codebase follows "Clean Architecture" principles. Logic is decoupled into distinct layers (`internal/` packages), making the code easy to read, test, and extend.

---

## 🚀 Features at a Glance

- **⚡ Blazing Fast**: Processes thousands of files per second using a concurrent worker pool architecture.
- **🛡️ Dry Run Mode**: Preview exactly what will happen without touching a single file. Perfect for testing.
- **🔄 Smart Duplicate Handling**: If `photo.jpg` already exists in the destination, the new file is automatically renamed to `photo_174567890.jpg` (timestamped) to prevent overwrites.
- **📂 Intelligent Categorization**: Recognizes over 60+ file extensions across 8 major categories.
- **📊 Detailed Analytics**: Provides a summary report with execution time, number of files moved, skipped, or failed, and total data size processed.
- **🖥️ Cross-Platform**: Compiles to a single static binary that runs on Windows, macOS, and Linux without dependencies.
- **🔍 Recursive Safety**: Deliberately does _not_ scan recursively by default to prevent destroying project structures, but handles flat directory organization perfectly.
- **📝 Verbose Logging**: Optional detailed logs to track every single file operation.

---

## 🏗️ Architecture & Design

We follow the standard Go project layout, ensuring the code is structured for production use.

### Directory Structure

```plaintext
file-organizer/
├── cmd/
│   └── file-organizer/
│       └── main.go           # The entry point. Parses flags and wires dependencies.
├── internal/
│   ├── app/
│   │   └── app.go            # Orchestration layer. Manages the worker pool and stats.
│   ├── classifier/
│   │   ├── classifier.go     # Domain logic. Maps extensions to categories.
│   │   └── classifier_test.go
│   └── fileops/
│       ├── mover.go          # Infrastructure layer. Handles OS file system calls.
│       └── mover_test.go
├── go.mod                    # Go module definition.
└── README.md                 # This documentation.
```

### Component Breakdown

1.  **`cmd/file-organizer`**: This is the "Main" package. It knows about the command-line arguments (flags) and responsible for instantiating the core components (`App`, `Classifier`, `Mover`) and connecting them together. It detects the Operating System context.

2.  **`internal/app`**: This is the "Application" layer. It contains the business orchestration. It doesn't know _how_ to move a file (that's `fileops`) or _what_ a file is (that's `classifier`), but it knows _when_ to do it. It manages the lifecycle of the application, the concurrency strategy (Worker Pool), and the collection of statistics.

3.  **`internal/classifier`**: This is a pure domain package. It contains the rules for organization. It maps strings (extensions) to other strings (Categories). It is side-effect free and easy to test.

4.  **`internal/fileops`**: This is the infrastructure layer. It wraps `os` and `path/filepath` calls. It encapsulates the "dirty work" of actually moving bytes on the disk, creating directories, and checking for file existence.

### Data Flow

1.  **User** runs the command with flags.
2.  **Main** parses flags and creates `App` config.
3.  **App** starts a `filesCh` channel and spawns `N` workers.
4.  **App** scans the directory and pushes file paths into `filesCh`.
5.  **Workers** pick up tasks:
    a. Call **Classifier** to get the category (e.g., ".jpg" -> "Images").
    b. Calculate destination path.
    c. Call **FileMover** to move the file.
6.  **Stats** are updated atomically.
7.  **App** prints the summary.

---

## 🔬 Deep Dive: Implementation Details

Let's explore the code logic module by module.

### The Entry Point (`cmd/file-organizer`)

Located in `cmd/file-organizer/main.go`.

This file is intentionally kept minimal. It performs **Dependency Injection**.

- It instantiates `classifier.NewClassifier()`.
- It instantiates `fileops.NewFileMover()`.
- It injects them into `app.App{}`.

This means the `App` struct depends on _instances_ rather than global helper functions, making it modular.

### Application Logic (`internal/app`)

Located in `internal/app/app.go`.

This is the brain of the operation.

- **Struct `App`**: Holds references to the dependencies.
- **Struct `Stats`**: Tracks metrics. Note the use of `int64` for fields like `FilesMoved`. This is crucial because we update these counters from multiple goroutines simultaneously, so we must use `sync/atomic` (e.g., `atomic.AddInt64`) to avoid race conditions.

### The Classifier Engine (`internal/classifier`)

Located in `internal/classifier/classifier.go`.

We use a `map[string]Category` for O(1) lookup time.

- **Normalization**: All extensions are converted to **lowercase** before lookup. This ensures `.JPG`, `.jpg`, and `.Jpg` are all treated as "Images".
- **Fallback**: If an extension is not in the map, it returns `Others`. This ensures no file is left behind, even if recognized.

### Safe File Operations (`internal/fileops`)

Located in `internal/fileops/mover.go`.

This module implements the **Strategy Pattern** for file movement.

- **`MoveFile` Function**:
  1.  Calculates the destination path.
  2.  Checks if the destination directory exists. If not, it creates it using `os.MkdirAll` (recursive mkdir).
  3.  **Collision Detection**: It runs `os.Stat(destPath)`. If the file exists, it triggers the renaming logic:
      ```go
      // Logic snippet
      timestamp := time.Now().UnixNano()
      newFileName := fmt.Sprintf("%s_%d%s", nameWithoutExt, timestamp, ext)
      ```
  4.  **Dry Run Check**: If `DryRun` is true, it logs the intent but _returns early_ without calling `os.Rename`.

---

## 🧵 Concurrency Model

One of the strongest features of this tool is its use of concurrency.

### Why Concurrency?

File I/O is often "blocking". If we organized files sequentially:

1.  Read File A.
2.  Wait for Move A to finish.
3.  Read File B.
4.  Wait for Move B...

If moving a large movie file takes 2 seconds, the program halts for 2 seconds.

### The Worker Pool Pattern

We implement a **Bounded Worker Pool**.

1.  **The Channel**: We create `filesCh := make(chan fileTask, 100)`. This is a buffered channel acting as a queue.
2.  **The Recursive Producer**: The main goroutine walks the directory and sends files into `filesCh`.
3.  **The Consumers (Workers)**: We spawn `runtime.NumCPU() * 2` persistent goroutines.
    - If your CPU has 8 cores, we spawn 16 workers.
    - Each worker sits in a loop: `for task := range filesCh`.
    - As soon as a worker finishes a move, it grabs the next file from the queue.

### Synchronization

- We use `sync.WaitGroup` to ensure the main program waits for all workers to finish emptying the queue before exiting.
- We call `close(filesCh)` only when we are done scanning the directory. This signals the workers that "no more work is coming".

---

## 🛡️ Safety Mechanisms

We treat your data with respect.

### Dry Run Mode

Enabled via `-dry-run`.

- This sets a boolean flag in the `FileMover`.
- All logic (classification, path calculation, collision checks) runs exactly as normal.
- Only the final `os.Rename` and `os.MkdirAll` calls are skipped.
- This gives you 100% confidence in the result before committing.

### Duplicate Collision Handling

A naive file organizer might overwrite `Resume.pdf` in the destination with the `Resume.pdf` from the source.
**We never overwrite.**

- Input: `Resume.pdf` -> Destination exists!
- Action: Rename to `Resume_1782312389.pdf`.
- Result: Both files are preserved.

### Atomic Stats

To prove the tool is working correctly, we track every success and failure.

- Using `sync/atomic` ensures that even if 16 threads try to increment "Files Moved" at the exact same nanosecond, the count remains accurate.

---

## 💾 Installation Guide

### Prerequisites

- **Go 1.20+**: You need the Go compiler installed. Type `go version` in your terminal to check.

### Building from Source

1.  **Download the Code**:
    Clone this repository or extract the zip to your workspace.

    ```bash
    cd path/to/13-cli-projects/file-organizer
    ```

2.  **Initialize dependencies** (if not already done):

    ```bash
    go mod tidy
    ```

3.  **Build the Binary**:
    This compiles the code into a standalone executable.

    **Windows (PowerShell)**:

    ```powershell
    go build -o file-organizer.exe ./cmd/file-organizer
    ```

    **Linux/macOS**:

    ```bash
    go build -o file-organizer ./cmd/file-organizer
    chmod +x file-organizer
    ```

4.  **Verify Installation**:
    ```bash
    ./file-organizer.exe -help
    ```

---

## 🎮 Usage Manual

### Command Line Flags

| Flag       | Type    | Default       | Description                                                    |
| :--------- | :------ | :------------ | :------------------------------------------------------------- |
| `-dir`     | String  | `.` (Current) | The source directory to organize.                              |
| `-dry-run` | Boolean | `false`       | Enable Preview Mode. Simulates actions without moving files.   |
| `-v`       | Boolean | `false`       | Enable Verbose Mode. Logs every file operation to the console. |
| `-help`    | Boolean | `false`       | Displays the help message and usage guide.                     |

### Examples

#### 1. The "Safe Check" (Recommended First Step)

Always run with dry-run first to see what will happen.

```bash
./file-organizer.exe -dir "C:\Users\Keshav\Downloads" -dry-run
```

_Output will show: "Would move X to Y" for all files._

#### 2. The "Clean Up"

Organize the default (current) directory.

```bash
./file-organizer.exe
```

#### 3. Organizing a Specific Folder

```bash
./file-organizer.exe -dir "D:\My Messy Desktop"
```

#### 4. Debugging Mode

If a file isn't moving, use verbose mode to see why.

```bash
./file-organizer.exe -dir "C:\Test" -v
```

---

## 🗂️ Supported File Types

The classifier currently supports the following mappings.

### 🖼️ Images

Folder: `Images`
Extensions: `.jpg`, `.jpeg`, `.png`, `.gif`, `.bmp`, `.svg`, `.webp`, `.tiff`, `.ico`, `.raw`, `.heic`

### 📄 Documents

Folder: `Documents`
Extensions: `.pdf`, `.doc`, `.docx`, `.txt`, `.rtf`, `.odt`, `.xls`, `.xlsx`, `.ppt`, `.pptx`, `.csv`, `.md`, `.epub`

### 🎵 Audio

Folder: `Audio`
Extensions: `.mp3`, `.wav`, `.flac`, `.aac`, `.ogg`, `.m4a`, `.wma`, `.aiff`

### 🎬 Video

Folder: `Video`
Extensions: `.mp4`, `.avi`, `.mkv`, `.mov`, `.wmv`, `.flv`, `.webm`, `.m4v`, `.mpeg`, `.mpg`

### 📦 Archives

Folder: `Archives`
Extensions: `.zip`, `.rar`, `.7z`, `.tar`, `.gz`, `.bz2`, `.xz`, `.iso`, `.tgz`

### 💻 Code

Folder: `Code`
Extensions: `.go`, `.py`, `.js`, `.html`, `.css`, `.java`, `.cpp`, `.c`, `.h`, `.ts`, `.json`, `.xml`, `.sql`, `.sh`, `.bat`, `.php`, `.rb`, `.pl`

### ⚙️ Executables

Folder: `Executables`
Extensions: `.exe`, `.msi`, `.apk`, `.app`, `.dmg`, `.deb`, `.rpm`, `.bin`, `.jar`

### 🔤 Fonts

Folder: `Fonts`
Extensions: `.ttf`, `.otf`, `.woff`, `.woff2`

_Any file not matching these categories is placed in an `Others` folder._

---

## 🧪 Testing & Quality Assurance

We trust our code because we test it.

To run the full test suite:

```bash
go test ./... -v
```

### What is tested?

1.  **Category Logic**: We assert that `image.JPG` is actually classified as an Image, and `unknown.foo` becomes "Others".
2.  **File Movement**: We create temporary files in a sandbox, run the mover, and verify they land in the right spot.
3.  **Collision Logic**: We create two matching files and ensure the second one gets renamed.
4.  **Dry Run Integrity**: We assert that in dry-run mode, files remain untouched.

---

## ❓ Troubleshooting

### "Access Denied" Errors

- **Cause**: The tool tries to move a file that is currently open in another program (e.g., a Word doc open in Word).
- **Solution**: Close all files in the target directory before running.
- **Note**: The tool handles this gracefully. It will log the error, skip that file, and continue organizing the rest. The failure count in the summary will increase.

### "Source path is not a directory"

- **Cause**: You provided a path to a file (e.g., `-dir ./photo.jpg`) instead of a folder.
- **Solution**: Ensure the path points to a folder.

### "Executables being moved"

- **Feature**: The tool is smart enough _not_ to move itself if the binary is inside the folder being organized.

---

## 🛤️ Future Roadmap

We are constantly improving. Here is what is planned for v2.0:

- [ ] **Custom Config**: Support for a `.organizer.yaml` file to define your own categories and extensions.
- [ ] **Recursive Mode**: Optional flag `-recursive` to organize subfolders (use with caution!).
- [ ] **Date-based Sorting**: Option to sort photos by Year/Month (e.g., `Images/2023/October`).
- [ ] **Undo Feature**: Keep a transaction log to reverse changes if needed.

---

## 🤝 Contributing

We love open source!

1.  **Fork** the repository.
2.  **Clone** your fork.
3.  **Create a Branch**: `git checkout -b feature/my-new-category`.
4.  **Add your feature** (e.g., add `.psd` to Images in `internal/classifier/classifier.go`).
5.  **Test**: Run `go test ./...`.
6.  **Push & PR**: Send us a Pull Request!

---

## 📄 License

This project is licensed under the **MIT License**. You are free to use, modify, and distribute this software.

---

_Generated for the Go 13-cli-projects suite._
_Author: Keshav & Antigravity (AI Agent)_
