package error

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/templates"
	"github.com/sirupsen/logrus"
)

/* Handlers */

// HTMLErrorHandler a html error handler
type HTMLErrorHandler struct {
	logger *httpErrorLogger
}

/* Builders */

type HTMLErrorHandlerBuilder struct {
	Logger *logrus.Logger
}

/* Create Functions */

func NewHTML() *HTMLErrorHandlerBuilder {
	return &HTMLErrorHandlerBuilder{log.StandardLogger()}
}

func (b *HTMLErrorHandlerBuilder) Create() *HTMLErrorHandler {
	logger := &httpErrorLogger{Logger: b.Logger}
	return &HTMLErrorHandler{
		logger: logger,
	}
}

func CreateHTML() *HTMLErrorHandler {
	return NewHTML().Create()
}

func (p *HTMLErrorHandler) sendError(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, dev *devInfo) {
	stacktrace := ""
	if dev != nil {
		stacktrace = string(dev.stacktrace)
	}
	func() {
		defer func() {
			// Error in html error handler -- fallback to plain error handler
			if err := recover(); err != nil {
				dev := getDevInfo(3) // FIXME get right offset
				Plain.panic(w, req, err, dev)
			}
		}()
		r := templates.Error.GetReadCloser(map[string]interface{}{
			"StatusCode":   statusCode,
			"ErrorMessage": statusMessage,
			"BackLink":     req.Referer(),
			"StackTrace":   stacktrace,
		})
		defer r.Close()
		w.WriteHeader(statusCode)
		if req.Method != "HEAD" {
			_, err := io.Copy(w, r)
			if err != nil {
				log.Info(err)
			}
		}
	}()
}

func (p *HTMLErrorHandler) error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage, logMessage string, dev *devInfo) {
	p.sendError(w, req, statusCode, statusMessage, dev)
	if 400 <= statusCode && statusCode < 500 { // Client Error
		p.logger.LogInfo(req, statusCode, logMessage, dev)
	} else {
		p.logger.LogError(req, statusCode, logMessage, dev)
	}
}

func (p *HTMLErrorHandler) Error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage, logMessage string) {
	dev := getDevInfo(1)
	p.error(w, req, statusCode, statusMessage, logMessage, dev)
}

func (p *HTMLErrorHandler) InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	dev := getDevInfo(1)
	p.error(w, req, http.StatusInternalServerError, "Internal Server Error", err.Error(), dev)
}

func (p *HTMLErrorHandler) panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *devInfo) {
	p.sendError(w, req, http.StatusInternalServerError, "Internal Server Error", dev)
	p.logger.LogPanic(req, http.StatusInternalServerError, fmt.Sprintf("%v", err), dev)
}
