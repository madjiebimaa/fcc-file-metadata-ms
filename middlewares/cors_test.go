package middlewares_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-file-metadata-ms/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	mid := middlewares.CORS()

	mid(c)

	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}
