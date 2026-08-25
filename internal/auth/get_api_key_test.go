package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPISuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key-987")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("error, got %v", err)
	}
	if key != "my-secret-key-987" {
		t.Fatalf("expected key %q, got %q", "my-secret_key-987", key)
	}
}

func TestGetAPIFailure(t *testing.T) {
	headers := http.Header{}

	key, err := GetAPIKey(headers)
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
	if key != "" {
		t.Errorf("expected empty key, got %q", key)
	}
}
