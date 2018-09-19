package resume

import (
	"net/http"

	"github.com/gorilla/mux"
	tmpldef "github.com/jsorrell/www.jacksorrell.com/templates/defs"
	"github.com/jsorrell/www.jacksorrell.com/web/pages"
)

var resumePage = pages.NewStaticPage(tmpldef.Resume,
	pages.PushStyle,
	pages.ServerPush("/js/contact.js"),
	pages.ServerPush("/img/beaker.svg"),
	pages.ServerPush("/img/briefcase.svg"),
	pages.ServerPush("/img/fork.svg"),
	pages.ServerPush("/img/location-pin.svg"),
	pages.ServerPush("/img/myface-nobg.png"),
	pages.ServerPush("/img/octocat.svg"),
	pages.ServerPush("/img/paperclip.svg"),
	pages.ServerPush("/img/person.svg"),
	pages.ServerPush("/img/terminal.svg"),
	pages.ServerPush("/img/Twitter_Social_Icon_Circle_Color.svg"),
	pages.ServerPush("/img/keybase_logo_official.svg"),
)

// RegisterRoutesTo registers routes to router
func RegisterRoutesTo(router *mux.Router) {
	router.Path("/resume/").Methods(http.MethodGet, http.MethodHead).Handler(resumePage)
}
