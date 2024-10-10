package builder

import (
	"html/template"
	"io"
	"os"

	"github.com/plutov/gitprint/api/pkg/files"
	"github.com/plutov/gitprint/api/pkg/log"
)

func GenerateHTML(w io.Writer, doc *Document, exportID string) error {
	logCtx := log.With("exportID", exportID)
	logCtx.Info("generating html")

	t, err := template.ParseFiles("./templates/base.html")
	if err != nil {
		logCtx.WithError(err).Error("failed to parse template")
		return err
	}

	if err := t.Execute(w, doc); err != nil {
		logCtx.WithError(err).Error("failed to execute template")
		return err
	}

	logCtx.Info("html generated")
	return nil
}

func GenerateAndSaveHTMLFile(doc *Document, exportID string) (string, error) {
	logCtx := log.With("exportID", exportID)
	logCtx.Info("saving html file")

	output := files.GetExportHTMLFile(exportID)

	if err := os.MkdirAll(output, 0755); err != nil {
		logCtx.WithError(err).Error("failed to create output directory")
		return "", err
	}
	o, err := os.Create(output)
	if err != nil {
		logCtx.WithError(err).Error("failed to create output file")
		return "", err
	}
	defer o.Close()

	if err := GenerateHTML(o, doc, exportID); err != nil {
		logCtx.WithError(err).Error("failed to generate html")
		return "", err
	}

	logCtx.Info("html file saved")
	return output, nil
}
