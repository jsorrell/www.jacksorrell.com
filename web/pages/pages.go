package pages

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/jsorrell/www.jacksorrell.com/log"
	myio "github.com/jsorrell/www.jacksorrell.com/utils/io"
)

// TODO do etags better -- not sure how to interact with templates in this way

// DefaultContentType is the content type that is used if none is specified.
const DefaultContentType = "text/html; charset=utf-8"

// PushStyle pushes "/css/style.css" to the client.
var PushStyle = ServerPush("/css/style.css")

var defaultStaticOptions = []PageOption{ContentType(DefaultContentType), PublicCache(86400)}
var defaultDynamicOptions = []PageOption{ContentType(DefaultContentType), NoCache}

type pageBase struct {
	headers map[string]string
}

/* Options */
// TODO Expand options

// PageOption is the type of options that is supplied to a page when created. These apply to all page types and later options given take precidence.
type PageOption func(*pageBase)

// ContentType is an option that sets the content type header of the page.
func ContentType(ct string) PageOption {
	return func(p *pageBase) {
		p.headers["Content-Type"] = ct
	}
}

// ServerPush is an option that adds a server push whenever the pages is requested.
func ServerPush(path string) PageOption {
	return func(p *pageBase) {
		p.headers["Link"] = appendPush(p.headers["Link"], path, detectServerPushType(path))
	}
}

// ServerPushAs is an option that adds a server push with the given "as" whenever the pages is requested.
func ServerPushAs(path, as string) PageOption {
	return func(p *pageBase) {
		p.headers["Link"] = appendPush(p.headers["Link"], path, as)
	}
}

// Header is an option that sets a header on the page.
func Header(key, value string) PageOption {
	return func(p *pageBase) {
		p.headers[key] = value
	}
}

// NoCache is an option that with completely disable caching of the page by the client.
func NoCache(p *pageBase) {
	p.headers["Cache-Control"] = "no-cache, no-store, must-revalidate"
	p.headers["Pragma"] = "no-cache"
	p.headers["Expires"] = "0"
}

// PublicCache is an option that with set the cache-control to public with the given max age.
func PublicCache(maxAge int) PageOption {
	return func(p *pageBase) {
		p.headers["Cache-Control"] = fmt.Sprintf("max-age=%d, public", maxAge)
	}
}

/* Static Page */

// StaticPageSource is an interface for a type that provides the ability to read the page content.
type StaticPageSource interface {
	GetReader() (myio.ReadSeekerCloser, string)
}

// StaticPage implements a static page for a website.
type StaticPage struct {
	pageBase
	s StaticPageSource
}

// NewStaticPage creates a new StaticPage with the given source and options. Later options override previous ones.
func NewStaticPage(source StaticPageSource, options ...PageOption) *StaticPage {
	page := &StaticPage{pageBase: pageBase{make(map[string]string, 0)}, s: source}

	for _, opt := range append(defaultStaticOptions, options...) {
		opt(&page.pageBase)
	}

	return page
}

func (page *StaticPage) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r, etag := page.s.GetReader()
	defer r.Close()
	w.Header().Set("Etag", etag)
	page.pageBase.serveHTTP(w, req, r)
}

/* Dynamic Page */

// DynamicPageSource is an interface for a type that provides the ability to read the page content when given an argument.
type DynamicPageSource interface {
	GetReader(vars interface{}) myio.ReadSeekerCloser
}

// DynamicPage implements a dynamic page for a website.
type DynamicPage struct {
	pageBase
	s DynamicPageSource
}

// NewDynamicPage creates a new DynamicPage with the given source and options. Later options override previous ones.
func NewDynamicPage(source DynamicPageSource, options ...PageOption) *DynamicPage {
	page := &DynamicPage{pageBase: pageBase{make(map[string]string, 0)}, s: source}

	for _, opt := range append(defaultDynamicOptions, options...) {
		opt(&page.pageBase)
	}

	return page
}

func (page *DynamicPage) ServeHTTP(w http.ResponseWriter, req *http.Request, args interface{}) {
	r := page.s.GetReader(args)
	defer r.Close()
	page.pageBase.serveHTTP(w, req, r)
}

/* Helpers */

func (page *pageBase) serveHTTP(w http.ResponseWriter, req *http.Request, reader io.ReadSeeker) {
	w.Header().Set("Content-Type", DefaultContentType)
	for k, v := range page.headers {
		w.Header().Set(k, v)
	}
	http.ServeContent(w, req, "", time.Unix(0, 0), reader)
}

// GenEtag FIXME this should be somewhere else
func GenEtag(data []byte) string {
	sum := md5.Sum(data)
	b64 := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(sum[:])
	return "\"" + b64 + "\""
}

func appendPush(spHeader, path, as string) string {
	if spHeader != "" {
		spHeader += ", "
	}
	return spHeader + fmt.Sprintf("<%s>; rel=preload; as=%s", path, as)
}

func detectServerPushType(path string) string {
	// TODO flesh this out or use 3rd party library
	switch filepath.Ext(path) {
	case ".css":
		return "style"
	case ".js":
		return "script"
	case ".ttf", ".woff":
		return "font"
	case ".svg", ".png", ".gif", ".ico", ".jpeg", ".jpg", ".tif", ".tiff":
		return "image"
	default:
		log.Warn("Could not detect push type for path " + path)
		return "embed"
	}
}
