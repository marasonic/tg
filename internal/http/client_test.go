package http
package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendPostRequestWithToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-token" {
			t.Errorf("Expected Authorization header 'Bearer test-token', got '%s'", authHeader)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := SendPostRequest(server.URL, "test-token", map[string]string{"foo": "bar"})
	if err != nil {
		t.Errorf("SendPostRequest failed: %v", err)
	}
}
