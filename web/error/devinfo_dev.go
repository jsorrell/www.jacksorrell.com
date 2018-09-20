// +build dev

package error

import (
	"runtime"
	"runtime/debug"
)

//go:noinline
func getDevInfo(offset int) *DevInfo {
	stacktrace := debug.Stack()
	_, file, line, ok := runtime.Caller(offset + 1)
	if !ok {
		file = "???"
		line = 0
	}

	return &DevInfo{stacktrace, file, line}
}
