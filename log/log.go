package log

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/negroni"

	"github.com/sirupsen/logrus"
)

type Logger *logrus.Logger

// HTTPLogger the logger for http requests that implements negroni.Logger
type HTTPRequestLogger struct {
	*logrus.Logger
}

func (l *HTTPRequestLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, req)

	res := rw.(negroni.ResponseWriter)

	l.WithFields(logrus.Fields{
		"dur":    formatDuration(time.Since(start)),
		"status": ColoredStatus(res.Status()),
	}).Debug(fmt.Sprintf("%s %s", req.Method, req.URL.Path))
}

func formatDuration(d time.Duration) string {
	// length total number of characters in the string
	const length = 7
	const decimals = 3
	const divisor = time.Millisecond
	const suffix = "ms"
	const durLen = length - len(suffix)
	durStr := fmt.Sprintf("%*.*f", durLen, decimals, float32(d)/float32(divisor))
	switch l := len(durStr); {
	case l == durLen:
		break
	case l-decimals < durLen || len(durStr)-decimals-1 == durLen:
		durStr = durStr[:durLen]
	case l-decimals == durLen:
		durStr = durStr[:durLen-1] + "."
	default:
		durStr = strings.Repeat("9", durLen)
	}
	return durStr + suffix
}

func ColoredStatus(status int) ColoredField {
	// var color color.SprintFunc
	switch status / 100 {
	case 1: // Informational
		return ColoredField{color.New(color.FgWhite), fmt.Sprintf("%-3d", status)} // Gray
	case 2: // Success
		return ColoredField{color.New(color.FgHiGreen), fmt.Sprintf("%-3d", status)}
	case 3: // Redirection
		return ColoredField{color.New(color.FgBlue), fmt.Sprintf("%-3d", status)}
	case 4: // Client Error
		return ColoredField{color.New(color.FgYellow), fmt.Sprintf("%-3d", status)}
	case 5: // Server Error
		return ColoredField{color.New(color.FgHiRed), fmt.Sprintf("%-3d", status)}
	default:
		return ColoredField{color.New(color.FgHiRed), fmt.Sprintf("%-3d", status)}
	}
}
