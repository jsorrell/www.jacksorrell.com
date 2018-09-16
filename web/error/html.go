package error

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/configloader"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/templates"
	"github.com/sirupsen/logrus"
)

/* Handlers */

// HTMLErrorHandler a html error handler
type HTMLErrorHandler struct {
	logger *httpErrorLogger
	dev    bool
}

/* Builders */

type HTMLErrorHandlerBuilder struct {
	Logger *logrus.Logger
	Dev    bool
}

/* Create Functions */

func NewHTML() *HTMLErrorHandlerBuilder {
	return &HTMLErrorHandlerBuilder{log.StandardLogger(), configloader.Config.Dev}
}

func (b *HTMLErrorHandlerBuilder) Create() *HTMLErrorHandler {
	logger := &httpErrorLogger{Logger: b.Logger, Dev: b.Dev}
	return &HTMLErrorHandler{
		logger: logger,
		dev:    b.Dev,
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
	tmpl, err := templates.GetTemplate("www")
	if err != nil {
		Plain.InternalServerError(w, req, err)
		return
	}

	w.WriteHeader(statusCode)
	tmpl.ExecuteTemplate(w, "error", map[string]interface{}{"StatusCode": statusCode, "ErrorMessage": statusMessage, "BackLink": req.Referer(), "StackTrace": stacktrace})
}

func (p *HTMLErrorHandler) error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, logMessage string, dev *devInfo) {
	p.sendError(w, req, statusCode, statusMessage, dev)
	if 400 <= statusCode && statusCode < 500 { // Client Error
		p.logger.LogInfo(req, statusCode, logMessage, dev)
	} else {
		p.logger.LogError(req, statusCode, logMessage, dev)
	}
}

func (p *HTMLErrorHandler) Error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, logMessage string) {
	var dev *devInfo
	if p.dev {
		dev = getDevInfo(1)
	}
	p.error(w, req, statusCode, statusMessage, logMessage, dev)
}

func (p *HTMLErrorHandler) InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	var dev *devInfo
	if p.dev {
		dev = getDevInfo(1)
	}
	p.error(w, req, http.StatusInternalServerError, "Internal Server Error", err.Error(), dev)
}

func (p *HTMLErrorHandler) panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *devInfo) {
	p.sendError(w, req, http.StatusInternalServerError, "Internal Server Error", dev)
	p.logger.LogPanic(req, http.StatusInternalServerError, fmt.Sprintf("%v", err), dev)
}

func (p *HTMLErrorHandler) isDev() bool {
	return p.dev
}
