package resume

import (
	"bytes"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jsorrell/www.jacksorrell.com/templates"
)

var bufPool *sync.Pool

func init() {
	bufPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 50000))
		},
	}
}

// RegisterRoutesTo registers routes to router
func RegisterRoutesTo(router *mux.Router) {
	router.Path("/resume/").Methods("GET", "HEAD").Handler(templates.Resume)
}
