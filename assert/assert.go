package assert

import (
	"github.com/Etuloser/learn-go-with-tests/pkg/ereflect"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper interface {
	Helper()
}

func Equal(t TestingT, expected interface{}, actual interface{}, msg ...string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if !ereflect.Equal(expected, actual) {
		t.Errorf("%s expected=%+v, actual=%+v", msg, expected, actual)
	}
}

func IsNotNil(t TestingT, actual interface{}, msg ...string) {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	if ereflect.IsNil(actual) {
		t.Errorf("%s expected not nil, but actual=%+v", msg, actual)
	}
}
