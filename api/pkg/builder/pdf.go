package builder

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/plutov/gitprint/api/pkg/files"
	"github.com/plutov/gitprint/api/pkg/log"
)

func GenerateAndSavePDFFile(htmlFile string, exportID string) (string, error) {
	logCtx := log.With("exportID", exportID, "html_file", htmlFile)
	logCtx.Info("generating pdf")

	output := files.GetExportPDFFile(exportID)

	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		logCtx.WithError(err).Error("failed to create output directory")
		return "", err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("index.html", "index.html")
	if err != nil {
		logCtx.WithError(err).Error("failed to create form file")
		return "", err
	}

	f, err := os.Open(htmlFile)
	if err != nil {
		logCtx.WithError(err).Error("failed to open file")
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(fw, f); err != nil {
		logCtx.WithError(err).Error("failed to copy file")
		return "", err
	}

	fields := map[string]string{
		// TODO: adjust waitDelay
		"waitDelay":       "1s",
		"marginTop":       "10px",
		"marginBottom":    "10px",
		"marginLeft":      "0",
		"marginRight":     "0",
		"printBackground": "true",
		"pageRanges":      "1-100",
	}
	for k, v := range fields {
		if err := w.WriteField(k, v); err != nil {
			logCtx.WithError(err).Error("failed to write field")
			return "", err
		}
	}
	w.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/forms/chromium/convert/html", os.Getenv("GOTENBERG_ADDR")), &b)
	if err != nil {
		logCtx.WithError(err).Error("failed to create request")
		return "", err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := httpClient.Do(req)
	if err != nil {
		logCtx.WithError(err).Error("failed to send request")
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		logCtx.With("status_code", res.StatusCode).Error("unexpected status code")
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	o, err := os.Create(output)
	if err != nil {
		logCtx.WithError(err).Error("failed to create output file")
		return "", err
	}

	defer o.Close()

	// Write the response body to file
	if _, err := io.Copy(o, res.Body); err != nil {
		logCtx.WithError(err).Error("failed to copy response body")
		return "", err
	}

	return output, nil
}
