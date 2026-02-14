package classifier

import (
	"testing"
)

func TestClassifier_Classify(t *testing.T) {
	c := NewClassifier()

	tests := []struct {
		filename string
		want     Category
	}{
		{"image.jpg", Images},
		{"photo.PNG", Images}, // Case insensitive
		{"document.pdf", Documents},
		{"music.mp3", Audio},
		{"movie.mp4", Video},
		{"archive.zip", Archives},
		{"script.py", Code},
		{"program.exe", Executables},
		{"font.ttf", Fonts},
		{"unknown.xyz", Others},
		{"no_extension", Others},
		{".gitignore", Others}, // dotfiles
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := c.Classify(tt.filename); got != tt.want {
				t.Errorf("Classify(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestNewClassifier(t *testing.T) {
	c := NewClassifier()
	if c == nil {
		t.Fatal("NewClassifier returned nil")
	}
	if len(c.extMap) == 0 {
		t.Error("NewClassifier created empty map")
	}
}
