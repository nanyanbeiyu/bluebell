package test

import (
	v1 "bluebell/api/v1"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCaretPostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, v1.CaretPostHandler)

	body := `{
		"title":"test",
		"content":"just a test",
		"community_id":"1"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	assert.Contains(t, w.Body.String(), "需要登录")
}
