package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, Go. I'm number 8, again. You said: %s", string(req))
}
