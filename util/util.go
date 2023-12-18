package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ResponseData struct {
	Data []string `json:"data"`
}

func ListFilesInFolder(folderPath string) ([]string, error) {
	var filenames []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileName := strings.SplitN(path, fmt.Sprintf("%s/", folderPath), 2)[1]
			filenames = append(filenames, fileName)
		}

		return nil
	})

	return filenames, err
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func FormatResponse(e error, statusCode int, data *ResponseData, timestamp time.Time) map[string]any {
	if e != nil {
		return gin.H{
			"error":      e.Error(),
			"statusCode": statusCode,
			"timestamp":  timestamp,
			"message":    "failure",
		}
	}

	return gin.H{
		"statusCode": statusCode,
		"timestamp":  timestamp,
		"response":   data,
		"message":    "success",
	}
}
