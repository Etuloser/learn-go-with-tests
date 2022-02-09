package ereflect

import (
	"bytes"
	"reflect"
)

func IsNil(actual interface{}) bool {
	if actual == nil {
		return true
	}
	v := reflect.ValueOf(actual)
	k := v.Kind()
	if k == reflect.Chan || k == reflect.Map || k == reflect.Ptr || k == reflect.Interface || k == reflect.Slice {
		return v.IsNil()
	}
	return false
}

func Equal(expected, actual interface{}) bool {
	if expected == nil {
		return IsNil(actual)
	}

	exp, ok := expected.([]byte)
	if ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}
		return bytes.Equal(exp, act)
	}

	return reflect.DeepEqual(expected, actual)
}
