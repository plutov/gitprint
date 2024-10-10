package files

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/plutov/gitprint/api/pkg/log"
)

const (
	publicInternalDir = "gitprint_public_internal"
)

func GetExportDir(exportID string) string {
	return filepath.Join(os.Getenv("GITHUB_REPOS_DIR"), exportID)
}

func GetExportHTMLFile(exportID string) string {
	return filepath.Join(os.Getenv("GITHUB_REPOS_DIR"), publicInternalDir, exportID) + ".html"
}

func GetExportPDFFile(exportID string) string {
	return filepath.Join(os.Getenv("GITHUB_REPOS_DIR"), publicInternalDir, exportID) + ".pdf"
}

func GenerateExportID() string {
	timestamp := time.Now().UnixNano()

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.WithError(err).Error("unable to generate random bytes")
		return fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(int(timestamp)))))
	}

	salt := base64.URLEncoding.EncodeToString(b)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(salt+strconv.Itoa(int(timestamp)))))
}

func ValidateExportID(exportID string) error {
	if len(exportID) != 32 {
		return fmt.Errorf("invalid export_id")
	}

	return nil
}
