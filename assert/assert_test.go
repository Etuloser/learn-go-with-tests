package assert

import (
	"fmt"
	"os"
	"testing"
)

type MockTestingT struct {
}

func (mtt MockTestingT) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
}

func TestEqual(t *testing.T) {
}
