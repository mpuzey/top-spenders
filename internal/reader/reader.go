package reader

import (
	"encoding/csv"
	"io"
)

// RecordProcessor is an interface that defines a method to process a single record from a CSV file
type RecordProcessor interface {
	Process(record []string) error
}

// ReadCSV reads CSV data and processes each line sequentially
func ReadCSV(reader io.Reader, processor RecordProcessor) error {
	csvReader := csv.NewReader(reader)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := processor.Process(record); err != nil {
			return err
		}
	}

	return nil
}
