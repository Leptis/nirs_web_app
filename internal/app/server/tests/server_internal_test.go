package server_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"nirs_web_app/internal/app/server"
	"testing"

)


func TestServer_HandleHello(t *testing.T) {
	s := server.New(server.NewConfig())
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	s.HandleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")


}