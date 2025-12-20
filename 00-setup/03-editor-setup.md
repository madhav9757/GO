# Editor Setup

This section covers setting up **VS Code** for Go development, including extensions, formatting, and debugging.

## Install VS Code
1. Download VS Code from: https://code.visualstudio.com/
2. Install it using default settings.

## Install Go Extension
1. Open **VS Code**
2. Go to **Extensions** (`Ctrl + Shift + X`)
3. Search for **Go** (by Google)
4. Click **Install**

## Install Required Tools
1. Open Command Palette: `Ctrl + Shift + P`
2. Run:
    ```bash
    Go: Install/Update Tools
    ```
3. Select all tools and install them.

## Formatting Go Code
- Go uses `gofmt` for formatting.
- To format a file:
- Right-click → **Format Document**
- Or use shortcut: `Shift + Alt + F`

VS Code formats Go code automatically on save.

## Debugging Go Programs
1. Open a Go file
2. Click on the **Run and Debug** tab
3. Click **Run**
4. You can:
- Set breakpoints
- Step through code
- Inspect variables

## Recommended Settings
Add this to VS Code settings for better experience:
```json
{
"editor.formatOnSave": true,
"go.useLanguageServer": true
}
```
## Summary

✅ VS Code + Go extension is the best setup
🛠 Auto-formatting via gofmt
🐞 Built-in debugging support
🚀 Ready for professional Go development