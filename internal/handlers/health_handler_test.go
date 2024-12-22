package handlers

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	testCases := []struct {
		name         string
		wantStatus   int
		wantResponse []byte
	}{
		{
			name:         "check health",
			wantStatus:   200,
			wantResponse: []byte("It's alive!\n"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			HealthHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			// Проверяем статус ответа
			if status := resp.StatusCode; status != tt.wantStatus {
				t.Errorf("HealthHandler returned status %v, want %v", status, tt.wantStatus)
			}

			// Проверяем тело ответа
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %v", err)
			}

			if !bytes.Equal(body, tt.wantResponse) {
				t.Errorf("HealthHandler returned body %q, want %q", body, tt.wantResponse)
			}
		})
	}
}
