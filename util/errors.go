package util

import (
	"errors"
	"fmt"
)

// Lista de erros personalizados
var (
	ErrInvalidURL      = errors.New("Invalid URL")
	ErrFileNotFound    = errors.New("File not found")
	ErrFetchFailed     = errors.New("Failed to fetch content")
	ErrUnsupportedMode = errors.New("Unsupported Mode")
)

// WrapError adiciona contexto a um erro
func WrapError(err error, context string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}
