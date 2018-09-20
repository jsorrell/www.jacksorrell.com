package error

import (
	"fmt"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/log"
)

// Plain is the standard PlainText error handler.
var Plain = NewHandler(PlainErrorPageGenerator{})

// HTML is the standard PlainText error handler.
var HTML = NewHandler(HTMLErrorPageGenerator{})

// ErrorPageGenerator is the interface that defines how error is sent to the client.
// TODO find a good name for this.
type ErrorPageGenerator interface {
	// SendError responds to the request with a page to display the error.
	SendError(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, dev *DevInfo)
}

// Handler is a error handler that defines how the error is logged and the client is sent the error.
type Handler struct {
	Generator ErrorPageGenerator
	Logger    HTTPErrorLogger
}

// NewHandler returns a new Handler using the standard logger.
func NewHandler(g ErrorPageGenerator) *Handler {
	return NewHandlerWithLogger(g, HTTPErrorLogger{log.GetStandardLogger()})
}

// NewHandlerWithLogger returns a new Handler using the given logger.
func NewHandlerWithLogger(g ErrorPageGenerator, logger HTTPErrorLogger) *Handler {
	return &Handler{g, logger}
}

func (h *Handler) error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage, logMessage string, dev *DevInfo) {
	h.Generator.SendError(w, req, statusCode, statusMessage, dev)
	if 400 <= statusCode && statusCode < 500 { // Client Error
		h.Logger.Error(req, statusCode, logMessage, dev)
	} else {
		h.Logger.Error(req, statusCode, logMessage, dev)
	}
}

// Error sends a error message to the client and logs the error.
func (h *Handler) Error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage, logMessage string) {
	dev := getDevInfo(1)
	h.error(w, req, statusCode, statusMessage, logMessage, dev)
}

// InternalServerError sends an InternalServerError error message to the client and logs the error.
func (h *Handler) InternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	dev := getDevInfo(1)
	h.error(w, req, http.StatusInternalServerError, "Internal Server Error", err.Error(), dev)
}

// panic is private because it should only be called by the panic handler. The user should just call panic().
func (h *Handler) panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *DevInfo) {
	h.Generator.SendError(w, req, http.StatusInternalServerError, "Internal Server Error", dev)
	h.Logger.Panic(req, http.StatusInternalServerError, fmt.Sprintf("%v", err), dev)
}

/***************/
/* 404 Handler */
/***************/

// NotFoundHandler a handler of Not Found errors that implements http.Handler.
type NotFoundHandler struct {
	h *Handler
}

// GetNotFoundHandler returns a new NotFoundHandler.
func (h *Handler) GetNotFoundHandler() NotFoundHandler {
	return NotFoundHandler{h}
}

func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.Error(w, r, http.StatusNotFound, "Not Found", fmt.Sprintf("Resource %s not found", r.URL.Path))
}

/***************/
/* 405 Handler */
/***************/

// MethodNotAllowedHandler is a handler of Method Not Allowed errors that implements http.Handler.
type MethodNotAllowedHandler struct {
	h *Handler
}

// GetMethodNotAllowedHandler returns a new MethodNotAllowedHandler.
func (h *Handler) GetMethodNotAllowedHandler() MethodNotAllowedHandler {
	return MethodNotAllowedHandler{h}
}

func (h MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.Error(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", fmt.Sprintf("Method %s not allowed for %s", r.Method, r.URL.Path))
}

/*****************/
/* Panic Handler */
/*****************/

// Recoverer is middleware that recovers from panics that occur in future handlers. It implements negroni.Handler.
type Recoverer struct {
	h *Handler
}

// GetRecoverer returns a new Recoverer.
func (h *Handler) GetRecoverer() Recoverer {
	return Recoverer{h}
}

func (rec Recoverer) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			dev := getDevInfo(3)
			rec.h.panic(w, r, err, dev)
		}
	}()
	next(w, r)
}
