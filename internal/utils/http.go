package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func EncodeWithStatus[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func EncodeWithError(w http.ResponseWriter, status int, msg string, err ...error) error {
	fmt.Printf("%d %s: %v\n", status, msg, err)
	return EncodeWithStatus(w, status, map[string]string{"error": msg})
}

func Encode[T any](w http.ResponseWriter, v T) error {
	return EncodeWithStatus(w, http.StatusOK, v)
}

func Decode[T any](r *http.Request) (T, error) {
	var v T
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil && err != io.EOF {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func GetBearerToken(r *http.Request) (string, bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		fmt.Println("no Authorization token found in header")
		return "", false
	}
	const prefix = "Bearer "
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		fmt.Println("prefix length mismatch")
		return "", false
	}

	return auth[len(prefix):], true
}
