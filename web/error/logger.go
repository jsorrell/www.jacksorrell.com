package error

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/sirupsen/logrus"
)

type httpErrorLogger struct {
	Logger *logrus.Logger
	Dev    bool
}

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

func (l *httpErrorLogger) getEntry(req *http.Request, statusCode int, dev *devInfo) *logrus.Entry {
	logger := l.Logger.WithField("status", log.ColoredStatus(statusCode))
	if dev != nil {
		logger = logger.WithField("cause", fmt.Sprintf("%s:%d", getDisplayFile(dev.file), dev.line))
	}

	return logger
}

func (l *httpErrorLogger) LogInfo(req *http.Request, statusCode int, logMessage string, dev *devInfo) {
	l.getEntry(req, statusCode, dev).Info(logMessage)
}

func (l *httpErrorLogger) LogError(req *http.Request, statusCode int, logMessage string, dev *devInfo) {
	l.getEntry(req, statusCode, dev).Error(logMessage)
}

func (l *httpErrorLogger) LogPanic(req *http.Request, statusCode int, logMessage string, dev *devInfo) {
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
