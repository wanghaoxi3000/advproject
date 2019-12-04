package test

import (
	"advancedproject/serializer"
	"advancedproject/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouterStatus(t *testing.T) {
	router := server.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/status/", nil)
	router.ServeHTTP(w, req)

	response := serializer.Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", response.Data)
}
