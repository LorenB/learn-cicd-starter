package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantKey    string
		wantErr    error
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey my-secret-key",
			wantKey:    "my-secret-key",
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("unex[ected error: %v", err)
				}
			}

			if got != tt.wantKey {
				t.Errorf("expected key %q, got %q", tt.wantKey, got)
			}
		})
	}
}
