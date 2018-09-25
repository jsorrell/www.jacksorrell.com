package static

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/urfave/negroni"

	"github.com/jsorrell/www.jacksorrell.com/data"
)

// Static the type of server's static file handler
type Static struct{}

var publicStatic = negroni.NewStatic(data.WebPublic)
var faviconStatic = negroni.NewStatic(data.Favicons)

func (Static) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Only serve files with extension statically
	if filepath.Ext(r.URL.Path) == "" {
		next(rw, r)
		return
	}

	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
	}

	rw.Header().Set("Cache-Control", "max-age="+getCacheTime(upath)+", public")

	if !strings.ContainsRune(upath[1:], '/') {
		// Files in root redirected to favicon folder
		faviconStatic.ServeHTTP(rw, r, next)
	} else {
		publicStatic.ServeHTTP(rw, r, next)
	}
}

// GetStatic returns the server's static file handler.
func GetStatic() Static {
	return Static{}
}

// getCacheTime defines the cache times for static files by extensions here.
func getCacheTime(filename string) string {
	switch filepath.Ext(filename) {
	case ".js":
		return "604800" // 7 days
	case ".css":
		return "604800" // 7 days
	case ".svg", ".png", ".gif", ".jpg", ".jpeg":
		return "2592000" // 30 days
	default:
		return "86400" // 1 day
	}
}
