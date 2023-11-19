package error

import (
	"fmt"
)

func NotFound(template string, args ...interface{}) error {
	message := fmt.Sprintf("not found: %s", template)

	return fmt.Errorf(message, args...)
}
