// +build nocache

package resume

import (
	"io"

	"github.com/jsorrell/www.jacksorrell.com/templates"
)

func writeResume(w io.Writer) error {
	resData, err := parseResumeData()
	if err != nil {
		return err
	}

	tmpl, err := templates.GetTemplate("www")
	if err != nil {
		return err
	}

	return tmpl.ExecuteTemplate(w, "resume", map[string]interface{}{"ResumeData": resData})
}
