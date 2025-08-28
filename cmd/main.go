package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

type RecordProcessor interface {
	Process(record []string) error
}

type TransactionsCSVProcessor struct {
	ProcessedRecords [][]string
	mu               sync.Mutex
}

func (m *TransactionsCSVProcessor) Process(record []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Copy to avoid race conditions
	// recordCopy := make([]string, len(record))
	// copy(recordCopy, record)
	// m.ProcessedRecords = append(m.ProcessedRecords, recordCopy)
	fmt.Printf("unimplemented")
	return nil
}

// ProcessCSV
func ProcessCSV(reader io.Reader, processor RecordProcessor) error {
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

func main() {
	csvName := os.Getenv("CSV_FILE")
	if len(csvName) == 0 {
		csvName = "sample-transactions.csv"
	}

	file, err := os.Open(csvName)
	if err != nil {
		fmt.Printf("Error opening CSV file \"%s\" :%s", csvName, err)
		return
	}
	defer file.Close()

	processor := &TransactionsCSVProcessor{}
	err = ProcessCSV(file, processor)
	if err != nil {
		fmt.Printf("Error processing file(/s) %v", err)
	}
}
