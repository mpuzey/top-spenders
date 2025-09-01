package mock

import (
	"errors"
	"sync"
)

// Mock processor for testing
type MockProcessor struct {
	ProcessedRecords [][]string
	ShouldError      bool
	mu               sync.Mutex
}

// Process appends the record to the records list on the mock processor
func (m *MockProcessor) Process(record []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.ShouldError {
		return errors.New("processing error")
	}

	m.ProcessedRecords = append(m.ProcessedRecords, record)
	return nil
}
