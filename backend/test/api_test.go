package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github-commit-viewer/pkg/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCommits(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/commits?username=testuser", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetCommitsByRepo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/commits/testrepo?username=testuser", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
