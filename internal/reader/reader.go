package csv

import (
	"encoding/csv"
	"io"
)

// RecordProcessor is an interface that defines a method to process a single record from a CSV file
type RecordProcessor interface {
	Process(record []string) error
}

// ReadCSV streams CSV io and processes each line in a memory constrained/ efficient manner by using a buffered channel
func ReadCSV(reader io.Reader, processor RecordProcessor) error {
	records := make(chan []string, 100)
	errors := make(chan error, 1)

	go func() {
		defer close(records)
		csvReader := csv.NewReader(reader)
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				errors <- err
				return
			}
			records <- record
		}
	}()

	for record := range records {
		if err := processor.Process(record); err != nil {
			return err
		}

		select {
		case err := <-errors:
			return err
		default:
		}
	}

	return nil
}
