package utils

import (
	"errors"
	"strings"
)

var BLOCK_LIST = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}

const REPLACEMENT = "****"

func ValidateChirp(chirp string) error {
	if chirp == "" {
		return errors.New("Missing chirp")
	}
	if len(chirp) > 140 {
		return errors.New("chirp too long")
	}
	return nil
}

func CleanChirp(chirp string) string {
	parts := strings.Split(chirp, " ")
	for i, p := range parts {
		for _, w := range BLOCK_LIST {
			if strings.EqualFold(p, w) {
				parts[i] = REPLACEMENT
				break
			}
		}
	}
	return strings.Join(parts, " ")
}
