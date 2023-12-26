package errors

import "fmt"

type LudesError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (le *LudesError) Error() string {
	return fmt.Sprintf("%d: %s", le.Status, le.Message)
}

// NewLudesError returns new LudesError.
func NewLudesError(status int, message string) *LudesError {
	return &LudesError{
		Status:  status,
		Message: message,
	}
}
