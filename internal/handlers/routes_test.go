package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	s := &Server{port: 8080}
	router := gin.Default()
	router.POST("/login", s.Login)

	body := []byte(`{"username": "testuser"}`)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.Code)
	}

	var result map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	if err != nil {
		t.Errorf("failed to parse JSON response: %v", err)
	}

	if result["token"] == "" {
		t.Errorf("expected token. empty")
	}
}
