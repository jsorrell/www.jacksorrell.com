package templates

import (
	"html/template"
	"strings"

	"github.com/jsorrell/www.jacksorrell.com/config"
	"github.com/jsorrell/www.jacksorrell.com/log"
	resumeData "github.com/jsorrell/www.jacksorrell.com/routes/resume/data"
	tmpl "github.com/jsorrell/www.jacksorrell.com/templates/templateexecuter"
)

var webGroup = tmpl.NewTemplateGroup(
	"www",
	template.FuncMap{
		"contactMaxLength": config.ContactMaxLength,
		"resumeGenerateStars": func(stars int) string {
			if stars < 0 {
				stars = 0
			}
			if stars > 5 {
				stars = 5
			}
			return strings.Repeat("★", stars) + strings.Repeat("☆", 5-stars)
		},
	},
)

var Contact = webGroup.NewStaticTemplate("contactPage", "contact/contact-page.gohtml", nil).AddServerPushes(
	tmpl.ServerPush{Path: "/css/style.css", As: "style"},
)

var Resume = webGroup.NewStaticTemplate("resume", "resume/resume.gohtml", func() interface{} {
	resData, err := resumeData.ParseResumeData()
	if err != nil {
		log.Panic(err)
	}

	return map[string]interface{}{"ResumeData": resData}
}).AddServerPushes(
	tmpl.ServerPush{Path: "/css/style.css", As: "style"},
	tmpl.ServerPush{Path: "/js/contact.js", As: "script"},
	tmpl.ServerPush{Path: "/img/beaker.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/briefcase.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/fork.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/location-pin.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/myface-nobg.png", As: "image"},
	tmpl.ServerPush{Path: "/img/octocat.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/paperclip.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/person.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/terminal.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/Twitter_Social_Icon_Circle_Color.svg", As: "image"},
	tmpl.ServerPush{Path: "/img/keybase_logo_official.svg", As: "image"},
)

var Error = webGroup.NewDynamicTemplate("error", "error/error.gohtml").AddServerPushes(
	tmpl.ServerPush{Path: "/css/style.css", As: "style"},
)

func init() {
	webGroup.NewIncludedTemplate("contactBox", "contact/contact-box.gohtml")
	webGroup.NewIncludedTemplate("floatingContact", "contact/floating-contact.gohtml")
	webGroup.NewIncludedTemplate("favicon", "includes/favicon.gohtml")
	webGroup.Init()
}
