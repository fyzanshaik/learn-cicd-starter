package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectedError  string
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key-123"},
			},
			expectedAPIKey: "test-api-key-123",
			expectedError:  "",
		},
		{
			name:           "missing authorization header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedError:  "no authorization header included",
		},
		{
			name: "malformed authorization header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-token"},
			},
			expectedAPIKey: "",
			expectedError:  "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			if apiKey != tt.expectedAPIKey {
				t.Errorf("expected API key %q, got %q", tt.expectedAPIKey, apiKey)
			}

			if tt.expectedError == "" && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if tt.expectedError != "" && (err == nil || err.Error() != tt.expectedError) {
				t.Errorf("expected error %q, got %v", tt.expectedError, err)
			}
		})
	}
}
