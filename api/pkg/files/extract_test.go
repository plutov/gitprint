package files

import (
	"os"
	"testing"
)

func TestExtractAndFilterFiles(t *testing.T) {
	tests := []struct {
		path     string
		isNilErr bool
		files    int
		exclude  string
	}{
		{"notfound.tar.gz", false, 0, ""},
		{"./testdata/formulosity-0.1.5.tar.gz", true, 86, "ui/src/styles/*"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			res, err := ExtractAndFilterFiles(tt.path, tt.exclude)
			if res != nil {
				os.RemoveAll(res.OutputDir)
			}
			if tt.isNilErr && err != nil {
				t.Errorf("expecting nil error, got %v", err)
			}
			if tt.isNilErr == true && res.Files != tt.files {
				t.Errorf("expecting %d files, got %d", tt.files, res.Files)
			}
		})
	}
}
