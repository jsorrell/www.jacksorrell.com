// +build nocache

package contact

import (
	"io"

	"github.com/jsorrell/www.jacksorrell.com/templates"
)

func writeContactPage(w io.Writer) error {
	tmpl, err := templates.GetTemplate("www")
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "contactPage", nil)
}
