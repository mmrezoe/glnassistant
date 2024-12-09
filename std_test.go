// example_test.go
package glnassistant

import (
	"testing"
)

func TestStd(t *testing.T) {
	Stderr("Test text for error")
	Stdout("test", "true")
	Stdout("test", "")
}
