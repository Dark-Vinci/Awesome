package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/dark-vinci/node-fullstack/rusty/demo/handler"
)

func main() {
	router := gin.Default()

	h := handler.New("uploads")

	// ENDPOINTS
	router.POST("/upload", h.UploadFile)
	router.GET("/download/:filename", h.DownloadHandler)
	router.GET("/uploads", h.GetFileList)

	fmt.Println("<<<<<<APPLICATION NOW RUNNING>>>>>>")
	router.Run(":8080")
}
