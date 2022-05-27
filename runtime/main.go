package runtime

import "runtime"

func GetFileName() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}
