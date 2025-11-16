package writer

import (
	"encoding/csv"
	"os"

	"github.com/mr-isik/go-fred-scraper/pkg/fred"
)

// CSVWriter is a writer for writing FRED data to a CSV file
type CSVWriter struct {
	filePath string
}

// NewCSVWriter creates a new CSV writer
func NewCSVWriter(filePath string) *CSVWriter {
	return &CSVWriter{
		filePath: filePath,
	}
}

// Write writes the given observations to a CSV file
func (w *CSVWriter) Write(observations *fred.ObservationsResponse) error {
	file, err := os.Create(w.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"date", "value"}); err != nil {
		return err
	}

	// Write data
	for _, obs := range observations.Observations {
		if err := writer.Write([]string{obs.Date, obs.Value}); err != nil {
			return err
		}
	}

	return nil
}
