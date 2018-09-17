package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsorrell/www.jacksorrell.com/data"
	weblogger "github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/routes/contact"
	"github.com/jsorrell/www.jacksorrell.com/routes/resume"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
	"github.com/urfave/negroni"
)

// Middleware
var (
	recovery  = &weberror.Recoverer{ErrorHandler: weberror.HTML}
	static    = negroni.NewStatic(data.WebPublic)
	webLogger = weblogger.StandardHTTPRequestLogger()
)

// Server a server
type Server struct {
	*negroni.Negroni
}

// NewServer creates a new Server
func NewServer() *Server {
	s := Server{negroni.New()}

	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = &weberror.NotFoundHandler{ErrorHandler: weberror.HTML}
	router.MethodNotAllowedHandler = &weberror.MethodNotAllowedHandler{ErrorHandler: weberror.HTML}

	// Middleware
	s.Use(recovery)
	s.Use(static)
	s.Use(webLogger)

	// Controllers
	resume.RegisterRoutesTo(router)
	contact.RegisterRoutesTo(router)

	// Redirect to resume page
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/resume/", http.StatusTemporaryRedirect)
	})

	s.UseHandler(router)

	return &s
}
