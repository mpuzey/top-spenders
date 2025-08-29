package integration

import (
	"os"
	"testing"

	"github.com/matt/top-spenders/internal/csv"
	"github.com/matt/top-spenders/internal/csv/mock"

	"github.com/stretchr/testify/assert"
)

func TestProcessCSV_Integration(t *testing.T) {
	// Create temp file
	tmpFile, err := os.CreateTemp("", "test*.csv")
	if err != nil {

	}

	defer os.Remove(tmpFile.Name())

	// Write test data
	csvContent := "name,age\nAlice,25\nBob,30"
	_, err = tmpFile.WriteString(csvContent)
	if err != nil {
		t.Error("")
	}
	tmpFile.Close()

	// Test with actual file
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Errorf("Error opening temp file :%s", err)
	}
	defer file.Close()

	processor := &mock.MockProcessor{}
	err = csv.ReadCSV(file, processor)

	if err != nil {

	}
	assert.Len(t, processor.ProcessedRecords, 3)
}
