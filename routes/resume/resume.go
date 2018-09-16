package resume

import (
	"net/http"

	"github.com/gorilla/mux"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
)

// RegisterRoutesTo registers routes to router
func RegisterRoutesTo(router *mux.Router) {
	router.HandleFunc("/resume/", showResume)
}

func showResume(res http.ResponseWriter, req *http.Request) {
	err := writeResume(res)
	if err != nil {
		weberror.HTML.InternalServerError(res, req, err)
	}
}
