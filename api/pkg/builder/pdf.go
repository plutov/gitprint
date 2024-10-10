package builder

import (
	"os"

	"github.com/plutov/gitprint/api/pkg/files"
	"github.com/plutov/gitprint/api/pkg/log"
)

func GenerateAndSavePDFFile(htmlFile string, exportID string) (string, error) {
	logCtx := log.With("exportID", exportID, "html_file", htmlFile)
	logCtx.Info("generating pdf")

	output := files.GetExportPDFFile(exportID)

	if err := os.MkdirAll(output, 0755); err != nil {
		logCtx.WithError(err).Error("failed to create output directory")
		return "", err
	}

	return output, nil
}
