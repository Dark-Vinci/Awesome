package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dark-vinci/node-fullstack/rusty/demo/util"
)

type Handler struct {
	uploadPath string
}

func New(path string) Handler {
	return Handler{
		uploadPath: fmt.Sprintf("../%s", path),
	}
}

func (h *Handler) DownloadHandler(c *gin.Context) {
	fileName := c.Param("filename")
	filePath := fmt.Sprintf("%s/%s", h.uploadPath, fileName)

	if !util.FileExists(filePath) {
		c.JSON(http.StatusNotFound,
			util.FormatResponse(util.FileNotFountError,
				http.StatusNotFound,
				nil,
				time.Now(),
			),
		)
		return
	}

	c.File(filePath)
	return
}

func (h *Handler) GetFileList(c *gin.Context) {
	filenames, err := util.ListFilesInFolder(h.uploadPath)
	if err != nil {
		c.JSON(http.StatusNotFound,
			util.FormatResponse(util.InternalServerError,
				http.StatusInternalServerError,
				nil,
				time.Now(),
			),
		)
		return
	}

	// CHECK IF THE LENGTH OF THE FILE IS 0
	if len(filenames) == 0 {
		c.JSON(http.StatusNotFound,
			util.FormatResponse(util.NoFileUploaded,
				http.StatusNotFound,
				nil,
				time.Now(),
			),
		)
		return
	}

	c.JSON(http.StatusOK,
		util.FormatResponse(nil,
			http.StatusOK, &util.ResponseData{Data: filenames},
			time.Now(),
		),
	)
}

func (h *Handler) UploadFile(c *gin.Context) {
	// EXTRACT FILE FROM FORM DATA
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest,
			util.FormatResponse(util.FileNotUploaded,
				http.StatusBadRequest,
				nil,
				time.Now(),
			),
		)
		return
	}

	// SAVE FILE TO UPLOADS FOLDER
	filePath := fmt.Sprintf("%s/%s", h.uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError,
			util.FormatResponse(err,
				http.StatusInternalServerError,
				nil,
				time.Now(),
			),
		)
		return
	}

	response := util.ResponseData{Data: []string{fmt.Sprintf("File %s uploaded successfully", file.Filename)}}

	c.JSON(http.StatusOK, util.FormatResponse(nil, http.StatusOK, &response, time.Now()))
}
