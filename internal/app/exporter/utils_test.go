package exporter_test

import (
	"log"
	"os"

	"github.com/SafetyCulture/iauditor-exporter/internal/app/exporter"
)

// getTemporaryJSONExporter creates a JSONExporter that writes to a temp folder
func getTemporaryJSONExporter() exporter.Exporter {
	dir, err := os.MkdirTemp("", "export")
	if err != nil {
		log.Fatal(err)
	}

	return exporter.NewJSONExporter(dir)
}
