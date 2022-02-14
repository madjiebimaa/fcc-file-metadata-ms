package handlers_test

import (
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-file-metadata-ms/handlers"
	"github.com/madjiebimaa/fcc-file-metadata-ms/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFile_Analysis(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return 200 status code when upload an file", func(t *testing.T) {
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()
			file, _ := helpers.GetImageFromFilePath("./../uploaded/prince.png")
			part, err := writer.CreateFormFile("file", "prince.png")
			assert.NoError(t, err)

			err = png.Encode(part, file)
			assert.NoError(t, err)
		}()

		req := httptest.NewRequest("POST", "/", pr)
		req.Header.Add("Content-Type", writer.FormDataContentType())

		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		fileHandler := handlers.NewFileHandler()
		r.POST("/", fileHandler.Analysis)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 status code when file is not exist", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/", nil)
		rec := httptest.NewRecorder()
		_, r := gin.CreateTestContext(rec)
		fileHandler := handlers.NewFileHandler()
		r.POST("/", fileHandler.Analysis)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
