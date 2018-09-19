package error

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/sirupsen/logrus"
)

// PlainErrorHandler a plain text error handler
type PlainErrorHandler struct {
	logger *httpErrorLogger
}

type PlainErrorHandlerBuilder struct {
	Logger *logrus.Logger
}

/* Create Functions */
func New() *PlainErrorHandlerBuilder {
	return &PlainErrorHandlerBuilder{log.StandardLogger()}
}

func (b *PlainErrorHandlerBuilder) Create() *PlainErrorHandler {
	logger := &httpErrorLogger{Logger: b.Logger}
	return &PlainErrorHandler{
		logger: logger,
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

	if req.Method == http.MethodHead {
		return
	}

	_, err := w.Write([]byte(statusMessage))
	if err != nil {
		log.Info(err)
	}
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
	dev := getDevInfo(1)
	p.error(w, req, statusCode, statusMessage, logMessage, dev)
}

func (p *PlainErrorHandler) InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	dev := getDevInfo(1)
	p.error(w, req, http.StatusInternalServerError, "Internal Server Error", err.Error(), dev)
}

func (p *PlainErrorHandler) panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *devInfo) {
	p.sendError(w, req, http.StatusInternalServerError, "Internal Server Error")
	p.logger.LogPanic(req, http.StatusInternalServerError, fmt.Sprintf("%v", err), dev)
}
