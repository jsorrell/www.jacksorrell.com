package resume

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jsorrell/www.jacksorrell.com/log"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
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
	router.Path("/resume/").Methods("GET", "HEAD").HandlerFunc(showResume)
}

func showResume(res http.ResponseWriter, req *http.Request) {
	buf := bufPool.Get().(*bytes.Buffer)

	err := writeResume(buf)
	if err != nil {
		weberror.HTML.InternalServerError(res, req, err)
	}

	if req.Method == "HEAD" {
		buf.Reset()
		bufPool.Put(buf)
		res.WriteHeader(200)
		return
	}

	_, err = res.Write(buf.Bytes())
	buf.Reset()
	bufPool.Put(buf)
	if err != nil {
		log.Info(err)
	}
}
