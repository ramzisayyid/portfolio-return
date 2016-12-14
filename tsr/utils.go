package tsr

import (
	"fmt"
)

// ArgumentError is raised when one of the arguments sent to a function is
// invalid
type ArgumentError struct {
	arg string
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("incorrect value for %s", e.arg)
}
