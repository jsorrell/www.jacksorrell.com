package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jsorrell/www.jacksorrell.com/data"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/routes/contact"
	"github.com/jsorrell/www.jacksorrell.com/routes/resume"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
	"github.com/urfave/negroni"
)

// Middleware
var (
	recovery  = &weberror.Recoverer{ErrorHandler: weberror.HTML}
	static    = negroni.NewStatic(data.WebPublic)
	webLogger = log.StandardHTTPRequestLogger()
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

	router.Path("/favicon.ico").Methods(http.MethodGet, http.MethodHead).HandlerFunc(serveFavicon)

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

// TODO move this somewhere else
func serveFavicon(rw http.ResponseWriter, r *http.Request) {
	file := "fav/favicon.ico"
	f, err := data.WebPublic.Open(file)
	if err != nil {
		log.Error("fav/favicon.ico missing")
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Error("can't read fav/favicon.ico")
		return
	}

	http.ServeContent(rw, r, file, fi.ModTime(), f)
}
