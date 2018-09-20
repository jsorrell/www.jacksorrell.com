package error

import (
	"fmt"
	"net/http"
)

// PlainErrorPageGenerator generates a plain text error page using the error template.
type PlainErrorPageGenerator struct{}

// SendError responds to the request with a page to display the error.
func (PlainErrorPageGenerator) SendError(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, dev *DevInfo) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	if req.Method != http.MethodHead {
		fmt.Fprintln(w, statusMessage)
	}
}
