package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-file-metadata-ms/models"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (f *FileHandler) Analysis(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		res := fmt.Sprintf("get form err: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": res,
		})
		return
	}

	dst := fmt.Sprintf("./../uploaded/upload_%s", file.Filename)
	c.SaveUploadedFile(file, dst)
	res := models.File{
		Name: file.Filename,
		Size: int(file.Size),
		Type: file.Header.Get("Content-Type"),
	}

	c.JSON(http.StatusOK, res)
}
