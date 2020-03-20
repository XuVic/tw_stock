package helper

import "fmt"

func NotNull(s string) error {
	return &NotNullError{s}
}

type NotNullError struct {
	nullValue string
}

func (e *NotNullError) Error() string {
	return fmt.Sprintf("%s should not be nil.", e.nullValue)
}
