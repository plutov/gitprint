package files

import (
	"path/filepath"
	"strings"
)

const (
	// 100 KB
	MaxFileSize = 100 * 1024
)

var (
	allowedFileExtensions = []string{
		// Programming languages
		".c", ".cpp", ".h", ".inc",
		".java",
		".py",
		".js", ".ts", ".jsx", ".tsx", ".dart", ".coffee", ".mjs",
		".rb", ".rake", ".gemspec",
		".php",
		".go",
		".rs",
		".swift", ".kt", ".kts", ".scala",
		".sql",
		".html", ".css", ".scss", ".sass", ".less",
		".r",
		".pl", ".pm6",
		".m", ".mm",
		".cs", ".vb",
		".asm",
		".f", ".for",
		".pas",
		".ps1",
		".adb",
		".lisp", ".cl",
		".scm",
		".pl",
		".hs",
		".erl",
		".ex",
		".clj",
		".groovy",
		".jl",
		".lua",
		".st",
		".cob", ".cpy",
		".sh", ".bash", ".zsh", ".fish",
		".zig",
		// Configuration files
		".yaml", ".yml",
		".json",
		".xml",
		".toml",
		".ini",
		".plist",
		// Templates
		".j2", ".tmpl", ".tpl",
		// Documentation
		".md",
		".txt",
		"license",
		"owners",
		// Build files
		"dockerfile",
		"makefile",
	}
	blacklistFileSuffixes = []string{
		".min.css",
		".min.js",
		".min.map",
		"package.json",
		"lock.json",
		".lock",
	}
	blacklistDirs = []string{
		"node_modules",
		"vendor",
		"dist",
		"public",
		"mocks",
		"third_party",
		"testdata",
		"data",
		"docs",
		"doc",
	}
	blacklistPrefixes = []string{
		// Everything that starts with a dot
		".",
		"_",
	}
)

var allowedFileExtensionsMap = make(map[string]struct{})
var blacklistDirsMap = make(map[string]struct{})

func init() {
	for _, ext := range allowedFileExtensions {
		allowedFileExtensionsMap[ext] = struct{}{}
	}
	for _, dir := range blacklistDirs {
		blacklistDirsMap[dir] = struct{}{}
	}
}

func IsAllowedFile(path string) bool {
	path = strings.ToLower(path)
	file := filepath.Base(path)
	ext := filepath.Ext(path)

	_, okExt := allowedFileExtensionsMap[ext]
	_, okFile := allowedFileExtensionsMap[file]
	if !okExt && !okFile {
		return false
	}

	for _, suf := range blacklistFileSuffixes {
		if strings.Contains(path, suf) {
			return false
		}
	}

	for _, prefix := range blacklistPrefixes {
		if strings.HasPrefix(path, prefix) {
			return false
		}
	}

	return true
}

func IsAllowedDir(path string) bool {
	parts := strings.Split(path, string(filepath.Separator))

	for _, dir := range parts {
		if _, ok := blacklistDirsMap[dir]; ok {
			return false
		}

		for _, prefix := range blacklistPrefixes {
			if strings.HasPrefix(dir, prefix) {
				return false
			}
		}
	}

	return true
}
