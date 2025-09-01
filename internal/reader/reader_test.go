package csv_test

import (
	"strings"
	"testing"

	csv "github.com/matt/top-spenders/internal/reader"
	"github.com/matt/top-spenders/internal/reader/mock"

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
			processor := &mock.MockProcessor{}

			err := csv.ReadCSV(reader, processor)

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
