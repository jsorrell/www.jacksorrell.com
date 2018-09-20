package web

import (
	"github.com/jsorrell/www.jacksorrell.com/log"
	weberror "github.com/jsorrell/www.jacksorrell.com/web/error"
	"github.com/jsorrell/www.jacksorrell.com/web/static"
	"github.com/urfave/negroni"
)

// Server a server
type Server struct {
	*negroni.Negroni
}

// NewServer creates a new Server
func NewServer() *Server {
	s := Server{negroni.New(
		// Middleware
		// Recovery first; Logger after Static so we do not log static file lookups.
		weberror.HTML.GetRecoverer(),
		static.GetStatic(),
		log.GetStandardHTTPRequestLogger(),
	)}
	s.UseHandler(createRouter())
	return &s
}
