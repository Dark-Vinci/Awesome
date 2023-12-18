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
	// CREATE TEMP FOLDER
	tempDir := "test_uploads"
	err := os.Mkdir(tempDir, os.ModePerm)
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// CREATE DUMMY FILES IN TEMP FOLDER
	files := []string{"file1.txt", "file2.txt", "subfolder/file3.txt"}
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", tempDir, file)
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		assert.NoError(t, err)
		_, err = os.Create(filePath)
		assert.NoError(t, err)
	}

	filenames, err := ListFilesInFolder(tempDir)
	assert.NoError(t, err)

	// CHECK IF EXPECTED FILENAME IS PRESENT
	expectedFilenames := []string{"file1.txt", "file2.txt", "subfolder/file3.txt"}
	assert.ElementsMatch(t, expectedFilenames, filenames)
}

func TestFileExists(t *testing.T) {
	// CREATE A TEMP FOLDER
	tempFile := "test_file.txt"
	_, err := os.Create(tempFile)
	assert.NoError(t, err)
	defer os.Remove(tempFile)

	// TEST WITH FILE THAT EXIT
	exists := FileExists(tempFile)
	assert.True(t, exists)

	// TEST WITH FILE THAT DOES NOT EXIT
	nonExistentFile := "non_existent_file.txt"
	exists = FileExists(nonExistentFile)
	assert.False(t, exists)
}

func TestFormatResponse(t *testing.T) {
	// TEST WITH ERROR
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

	// TEST WITHOUT ERROR
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
