package error

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/jsorrell/www.jacksorrell.com/log"
)

// HTTPErrorLogger formats and logs HTTP errors to the given log.
type HTTPErrorLogger struct {
	logger *log.Logger
}

// getDisplayFile formats a filename to:
// If in www.jacksorrell.com/ -> the absolute path below www.jacksorrell.com
// If in $GOPATH/src/ -> the absolute path below src.
// Otherwise, just the filename.
func getDisplayFile(file string) string {
	if splits := strings.Split(file, "www.jacksorrell.com/"); len(splits) > 1 {
		return splits[len(splits)-1]
	}

	if gopath, exists := os.LookupEnv("GOPATH"); exists {
		if splits := strings.Split(file, filepath.Join(gopath, "src")); len(splits) > 1 {
			return splits[len(splits)-1]
		}
	}

	return filepath.Base(file)
}

func (l *HTTPErrorLogger) getEntry(req *http.Request, statusCode int, dev *DevInfo) *logrus.Entry {
	logger := l.logger.WithField("status", log.ColoredStatus(statusCode))
	if dev != nil {
		logger = logger.WithField("cause", fmt.Sprintf("%s:%d", getDisplayFile(dev.file), dev.line))
	}

	return logger
}

// Info logs at INFO level.
func (l *HTTPErrorLogger) Info(req *http.Request, statusCode int, logMessage string, dev *DevInfo) {
	l.getEntry(req, statusCode, dev).Info(logMessage)
}

// Error logs at ERROR level.
func (l *HTTPErrorLogger) Error(req *http.Request, statusCode int, logMessage string, dev *DevInfo) {
	l.getEntry(req, statusCode, dev).Error(logMessage)
}

// Panic logs at PANIC level.
// Ignores any panics the logger calls.
func (l *HTTPErrorLogger) Panic(req *http.Request, statusCode int, logMessage string, dev *DevInfo) {
	entry := l.getEntry(req, statusCode, dev)
	if dev != nil {
		entry = entry.WithField("stacktrace", dev.stacktrace)
	}
	func() {
		defer func() {
			recover()
		}()
		entry.Panic(logMessage)
	}()
}
