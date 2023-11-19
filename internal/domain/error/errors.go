package error

import (
	"fmt"
)

func NotFound(template string, args ...interface{}) error {
	message := fmt.Sprintf("not found: %s", template)

	return fmt.Errorf(message, args...)
}

func InvalidInput(template string, args ...interface{}) error {
	message := fmt.Sprintf("invalid input: %s", template)

	return fmt.Errorf(message, args...)
}

func Unavailable(template string, args ...interface{}) error {
	message := fmt.Sprintf("unavailable: %s", template)

	return fmt.Errorf(message, args...)
}
