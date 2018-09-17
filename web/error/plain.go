package error

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/configloader"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/sirupsen/logrus"
)

// PlainErrorHandler a plain text error handler
type PlainErrorHandler struct {
	logger *httpErrorLogger
	dev    bool
}

type PlainErrorHandlerBuilder struct {
	Logger *logrus.Logger
	Dev    bool
}

/* Create Functions */
func New() *PlainErrorHandlerBuilder {
	return &PlainErrorHandlerBuilder{log.StandardLogger(), configloader.Config.Dev}
}

func (b *PlainErrorHandlerBuilder) Create() *PlainErrorHandler {
	logger := &httpErrorLogger{Logger: b.Logger, Dev: b.Dev}
	return &PlainErrorHandler{
		logger: logger,
		dev:    b.Dev,
	}
}

func CreatePlain() *PlainErrorHandler {
	return New().Create()
}

func (p *PlainErrorHandler) sendError(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string) {
	w.WriteHeader(statusCode)
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}

	if req.Method == "HEAD" {
		return
	}

	w.Write([]byte(statusMessage))
}

func (p *PlainErrorHandler) error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, logMessage string, dev *devInfo) {
	p.sendError(w, req, statusCode, statusMessage)
	if 400 <= statusCode && statusCode < 500 { // Client Error
		p.logger.LogInfo(req, statusCode, logMessage, dev)
	} else {
		p.logger.LogError(req, statusCode, logMessage, dev)
	}
}

func (p *PlainErrorHandler) Error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, logMessage string) {
	var dev *devInfo
	if p.dev {
		dev = getDevInfo(1)
	}
	p.error(w, req, statusCode, statusMessage, logMessage, dev)
}

func (p *PlainErrorHandler) InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	var dev *devInfo
	if p.dev {
		dev = getDevInfo(1)
	}
	p.error(w, req, http.StatusInternalServerError, "Internal Server Error", err.Error(), dev)
}

func (p *PlainErrorHandler) panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *devInfo) {
	p.sendError(w, req, http.StatusInternalServerError, "Internal Server Error")
	p.logger.LogPanic(req, http.StatusInternalServerError, fmt.Sprintf("%v", err), dev)
}

func (p *PlainErrorHandler) isDev() bool {
	return p.dev
}
