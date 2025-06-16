package main

import (
	"errors"
	"fmt"
	"os"
)

type ValidationError struct {
	Field string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
func main ()  {
}

err, _ := &ValidationError{Field: "email", Message: "invalid format"}

originalErr := errors.New("database connection failed")
wrappedErr := fmt.Errorf("query failed: %w", originalErr)

// Later, you can unwrap:
fmt.Println(errors.Unwrap(wrappedErr))

var ErrNotFound = errors.New("item not found")

func FindItem(id string) (*Item, error) {
	// ...
	return nil, ErrNotFound
}

// Usage:
if err == ErrNotFound {
// Handle not found case
}


func processFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", name, err)
	}
	defer f.Close()

	// Process file...
	return nil
}