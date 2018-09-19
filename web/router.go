package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsorrell/www.jacksorrell.com/routes/contact"
	"github.com/jsorrell/www.jacksorrell.com/routes/resume"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
)

func createRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Controllers
	resume.RegisterRoutesTo(router)
	contact.RegisterRoutesTo(router)

	// Redirect to resume page
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/resume/", http.StatusTemporaryRedirect)
	})

	// Error Handlers
	router.NotFoundHandler = weberror.HTML.GetNotFoundHandler()
	router.MethodNotAllowedHandler = weberror.HTML.GetMethodNotAllowedHandler()
	return router
}
