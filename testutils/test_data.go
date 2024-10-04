package testutils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Test data for generic loading of JSON data files.
type TestData[T any] struct {
	Input T `json:"input"`
}

// Create a new instance of `TestData` with the given input data.
func NewTestData[T any](defaultInput T) TestData[T] {
	path, exists := os.LookupEnv("TEST_DATA_PATH")
	if exists {
		file, err := os.ReadFile(filepath.Clean(path))
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}

		var testData TestData[T]
		err = json.Unmarshal(file, &testData)
		if err != nil {
			log.Fatalf("Failed to decode JSON: %v", err)
		}
		return testData
	}
	return TestData[T]{Input: defaultInput}
}
