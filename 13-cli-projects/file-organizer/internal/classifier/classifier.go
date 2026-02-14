package classifier

import (
	"path/filepath"
	"strings"
)

// Category represents a file category
type Category string

const (
	Images      Category = "Images"
	Documents   Category = "Documents"
	Audio       Category = "Audio"
	Video       Category = "Video"
	Archives    Category = "Archives"
	Code        Category = "Code"
	Executables Category = "Executables"
	Fonts       Category = "Fonts"
	Others      Category = "Others"
)

// DefaultMappings holds the default extension to category mappings
var DefaultMappings = map[string][]string{
	string(Images):      {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg", ".webp", ".tiff", ".ico", ".raw", ".heic"},
	string(Documents):   {".pdf", ".doc", ".docx", ".txt", ".rtf", ".odt", ".xls", ".xlsx", ".ppt", ".pptx", ".csv", ".md", ".epub"},
	string(Audio):       {".mp3", ".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".aiff"},
	string(Video):       {".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".webm", ".m4v", ".mpeg", ".mpg"},
	string(Archives):    {".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".xz", ".iso", ".tgz"},
	string(Code):        {".go", ".py", ".js", ".html", ".css", ".java", ".cpp", ".c", ".h", ".ts", ".json", ".xml", ".sql", ".sh", ".bat", ".php", ".rb", ".pl"},
	string(Executables): {".exe", ".msi", ".apk", ".app", ".dmg", ".deb", ".rpm", ".bin", ".jar"},
	string(Fonts):       {".ttf", ".otf", ".woff", ".woff2"},
}

// Classifier handles file categorization logic
type Classifier struct {
	extMap map[string]Category
}

// NewClassifier creates a new Classifier with default mappings
func NewClassifier() *Classifier {
	extMap := make(map[string]Category)
	for cat, extensions := range DefaultMappings {
		for _, ext := range extensions {
			extMap[strings.ToLower(ext)] = Category(cat)
		}
	}
	return &Classifier{extMap: extMap}
}

// Classify returns the category for a given filename
func (c *Classifier) Classify(filename string) Category {
	ext := strings.ToLower(filepath.Ext(filename))
	if cat, found := c.extMap[ext]; found {
		return cat
	}
	return Others
}
