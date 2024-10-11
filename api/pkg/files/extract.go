package files

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/plutov/gitprint/api/pkg/log"
)

type ExtractAndFilterResult struct {
	Files     int    `json:"files"`
	OutputDir string `json:"-"`
	ExportID  string `json:"exportID"`
}

func ExtractAndFilterFiles(path string, exclude string) (*ExtractAndFilterResult, error) {
	logCtx := log.With("path", path)
	logCtx.Info("extracting and filtering files")

	r, err := os.Open(path)
	if err != nil {
		logCtx.WithError(err).Error("failed to open file")
		return nil, err
	}
	defer r.Close()

	// delete the archive file after processing, but do not remove testdata files
	if !strings.HasPrefix(path, "./testdata/") {
		defer os.RemoveAll(path)
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		logCtx.WithError(err).Error("failed to create gzip reader")
		return nil, err
	}
	defer gzr.Close()

	exportID := GenerateExportID()
	res := &ExtractAndFilterResult{
		OutputDir: GetExportDir(exportID),
		ExportID:  exportID,
	}

	excludePatterns := getExcludePatterns(exclude)

	tr := tar.NewReader(gzr)
	for {
		header, err := tr.Next()

		switch {
		// if no more files are found return
		case err == io.EOF:
			logCtx.With("res", *res).Info("extracted and filtered files")
			if res.Files == 0 {
				return nil, fmt.Errorf("no files found in the archive")
			}
			return res, nil

		// return any other error
		case err != nil:
			logCtx.WithError(err).Error("failed to read tar header")
			return nil, err

		// if the header is nil, just skip it
		case header == nil:
			continue
		}

		// remove root folder name but keep the hierarchy
		// eg. plutov-formulosity-xx123/README.md -> README.md
		// // eg. plutov-formulosity-xx123/src/main.go -> src/main.go
		parts := strings.Split(header.Name, string(filepath.Separator))
		if len(parts) > 0 {
			header.Name = strings.Join(parts[1:], "/")
		}
		target := filepath.Join(res.OutputDir, header.Name)

		// exclude patterns
		isExcluded := false
		for _, pattern := range excludePatterns {
			isExcluded, _ = regexp.MatchString(wildCardToRegexp(pattern), header.Name)
			if isExcluded {
				break
			}
		}
		if isExcluded {
			continue
		}

		// check the file type
		switch header.Typeflag {
		case tar.TypeReg:
			// skip empty and big files
			if header.Size == 0 || header.Size > MaxFileSize {
				continue
			}
			// skip blacklisted dirs
			headerDir := filepath.Dir(header.Name)
			if headerDir != "." && !IsAllowedDir(headerDir) {
				continue
			}

			if !IsAllowedFile(header.Name) {
				continue
			}

			targetDir := filepath.Dir(target)
			if err := os.MkdirAll(targetDir, 0755); err != nil {
				logCtx.WithError(err).Error("failed to create directory")
				return nil, err
			}

			f, err := os.Create(target)
			if err != nil {
				logCtx.WithError(err).Error("failed to create file")
				return nil, err
			}

			if _, err := io.Copy(f, tr); err != nil {
				f.Close()
				logCtx.WithError(err).Error("failed to copy file contents")
				return nil, err
			}

			f.Close()
			res.Files++
		}
	}
}

func getExcludePatterns(exclude string) []string {
	if exclude == "" {
		return nil
	}

	parts := strings.Split(exclude, ",")
	filtered := []string{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			filtered = append(filtered, part)
		}
	}

	return filtered
}

func wildCardToRegexp(pattern string) string {
	var result strings.Builder
	for i, literal := range strings.Split(pattern, "*") {

		// Replace * with .*
		if i > 0 {
			result.WriteString(".*")
		}

		// Quote any regular expression meta characters in the
		// literal text.
		result.WriteString(regexp.QuoteMeta(literal))
	}
	return result.String()
}
