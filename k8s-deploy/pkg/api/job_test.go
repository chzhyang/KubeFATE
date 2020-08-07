package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJob(t *testing.T) {
	router := initGinEngine()
	var w *httptest.ResponseRecorder
	url := "/job/"
	w = Get(url, router)
	assert.Equal(t, 200, w.Code)
}
