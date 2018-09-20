package log

import (
	"github.com/sirupsen/logrus"
)

// Logger is the type of logger for the app.
type Logger struct {
	logrus.Logger
}
