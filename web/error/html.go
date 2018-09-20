package error

import (
	"io"
	"net/http"

	"github.com/jsorrell/www.jacksorrell.com/log"
	tmpldefs "github.com/jsorrell/www.jacksorrell.com/templates/defs"
)

/* Handlers */

// HTMLErrorPageGenerator generates a HTML error page using the error template.
type HTMLErrorPageGenerator struct{}

// SendError responds to the request with a page to display the error.
func (HTMLErrorPageGenerator) SendError(w http.ResponseWriter, req *http.Request, statusCode int, statusMessage string, dev *DevInfo) {
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
		r := tmpldefs.Error.GetReader(map[string]interface{}{
			"StatusCode":   statusCode,
			"ErrorMessage": statusMessage,
			"BackLink":     req.Referer(),
			"StackTrace":   stacktrace,
		})
		defer r.Close()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(statusCode)
		if req.Method != http.MethodHead {
			_, err := io.Copy(w, r)
			if err != nil {
				log.Info(err)
			}
		}
	}()
}
