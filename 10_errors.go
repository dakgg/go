package main

import (
	"errors"
	"fmt"
)

// ValidationError is returned when input validation fails.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s — %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

func errorHandling() {
	fmt.Println("\n=== 에러 처리 ===")

	for _, age := range []int{25, -1, 200} {
		if err := validateAge(age); err != nil {
			var ve *ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("  field=%s msg=%s\n", ve.Field, ve.Message)
			}
		} else {
			fmt.Printf("  age %d is valid\n", age)
		}
	}
}
