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

func (m *MockProcessor) Process(record []string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.ShouldError {
		return errors.New("processing error")
	}

	// // Copy to avoid race conditions
	// recordCopy := make([]string, len(record))
	// copy(recordCopy, record)
	// m.ProcessedRecords = append(m.ProcessedRecords, recordCopy)
	m.ProcessedRecords = append(m.ProcessedRecords, record)
	return nil
}
