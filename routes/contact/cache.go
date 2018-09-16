// +build !nocache

package contact

import (
	"io"
	"github.com/jsorrell/www.jacksorrell.com/templates"
	"github.com/jsorrell/www.jacksorrell.com/log"
	"bytes"
)

var _contactPageHTML []byte

func init() {
	tmpl, err := templates.GetTemplate("www")
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0, 5000))
	err = tmpl.ExecuteTemplate(buf, "contactPage", nil)
	if err != nil {
		log.Fatal(err)
	}
	_contactPageHTML = buf.Bytes()
}

func writeContactPage(w io.Writer) error {
	_, err := w.Write(_contactPageHTML)
	return err
}
