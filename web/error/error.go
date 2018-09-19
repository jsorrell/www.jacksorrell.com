package error

import (
	"fmt"
	"net/http"
)

var Plain = CreatePlain()

var HTML = CreateHTML()

type ErrorHandler interface {
	Error(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, logMessage string)
	InternalServerError(w http.ResponseWriter, req *http.Request, err error)
	panic(w http.ResponseWriter, req *http.Request, err interface{}, dev *devInfo)
}

/***************/
/* 404 Handler */
/***************/

// NotFoundHandler a handler of Not Found errors that implements http.Handler
type NotFoundHandler struct {
	h ErrorHandler
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.Error(w, r, http.StatusNotFound, "Not Found", fmt.Sprintf("Resource %s not found", r.URL.Path))
}

func (p *HTMLErrorHandler) GetNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{p}
}

/***************/
/* 405 Handler */
/***************/

// MethodNotAllowedHandler a handler of Method Not Allowed errors that implements http.Handler
type MethodNotAllowedHandler struct {
	h ErrorHandler
}

func (h *MethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.Error(w, r, http.StatusMethodNotAllowed, "Method Not Allowed", fmt.Sprintf("Method %s not allowed for %s", r.Method, r.URL.Path))
}

func (p *HTMLErrorHandler) GetMethodNotAllowedHandler() *MethodNotAllowedHandler {
	return &MethodNotAllowedHandler{p}
}

/*****************/
/* Panic Handler */
/*****************/

// Recoverer middleware that recovers from panics that occur in future handlers. Implements negroni.Handler
type Recoverer struct {
	h ErrorHandler
}

func (rec *Recoverer) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			dev := getDevInfo(3)
			rec.h.panic(w, r, err, dev)
		}
	}()
	next(w, r)
}

func (p *HTMLErrorHandler) GetRecoverer() *Recoverer {
	return &Recoverer{p}
}
