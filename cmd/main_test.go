package main

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCSV(t *testing.T) {
	tests := []struct {
		name        string
		csvContent  string
		expectError bool
		wantRecords [][]string
	}{
		{
			name:       "valid CSV",
			csvContent: "name,age\nJohn,25\nJane,30",
			wantRecords: [][]string{
				{"name", "age"},
				{"John", "25"},
				{"Jane", "30"},
			},
		},
		{
			name:        "malformed CSV",
			csvContent:  "name,age\nJohn,25,extra",
			expectError: true,
		},
		{
			name:        "empty CSV",
			csvContent:  "",
			wantRecords: [][]string(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.csvContent)
			processor := &MockProcessor{}

			err := ProcessCSV(reader, processor)

			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectError {
				assert.Equal(t, processor.ProcessedRecords, tt.wantRecords)
			}
		})
	}
}

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
