package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListFilesInFolder(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := "test_uploads"
	err := os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create some dummy files in the temporary directory
	files := []string{"file1.txt", "file2.txt", "subfolder/file3.txt"}
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", tempDir, file)
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		assert.NoError(t, err)
		_, err = os.Create(filePath)
		assert.NoError(t, err)
	}

	// Test the ListFilesInFolder function
	filenames, err := ListFilesInFolder(tempDir)
	assert.NoError(t, err)

	// Check if the expected filenames are present
	expectedFilenames := []string{"file1.txt", "file2.txt", "subfolder/file3.txt"}
	assert.ElementsMatch(t, expectedFilenames, filenames)
}

func TestFileExists(t *testing.T) {
	// Create a temporary file for testing
	tempFile := "test_file.txt"
	_, err := os.Create(tempFile)
	assert.NoError(t, err)
	defer os.Remove(tempFile)

	// Test if the file exists
	exists := FileExists(tempFile)
	assert.True(t, exists)

	// Test with a non-existent file
	nonExistentFile := "non_existent_file.txt"
	exists = FileExists(nonExistentFile)
	assert.False(t, exists)
}

func TestFormatResponse(t *testing.T) {
	// Test with an error
	err := errors.New("test error")
	statusCode := 500
	data := &ResponseData{Data: []string{"some data"}}
	ti := time.Now()

	result := FormatResponse(err, statusCode, data, ti)

	expectedError := map[string]interface{}{
		"error":      err.Error(),
		"statusCode": statusCode,
		"timestamp":  ti,
		"message":    "failure",
	}

	assert.Equal(t, expectedError, result)

	// Test without an error
	err = nil
	result = FormatResponse(err, statusCode, data, ti)

	expectedSuccess := map[string]interface{}{
		"statusCode": statusCode,
		"timestamp":  ti,
		"response":   data,
		"message":    "success",
	}

	assert.Equal(t, expectedSuccess, result)
}
