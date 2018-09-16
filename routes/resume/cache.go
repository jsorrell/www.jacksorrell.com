// +build !nocache

package resume

import (
	"bytes"
	"io"

	"github.com/jsorrell/www.jacksorrell.com/log"
	"github.com/jsorrell/www.jacksorrell.com/templates"
)

var _resumeHTML []byte

func init() {
	resData, err := parseResumeData()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := templates.GetTemplate("www")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(make([]byte, 0, 50000))
	err = tmpl.ExecuteTemplate(buf, "resume", map[string]interface{}{"ResumeData": resData})
	if err != nil {
		log.Fatal(err)
	}
	_resumeHTML = buf.Bytes()
}

func writeResume(w io.Writer) error {
	_, err := w.Write(_resumeHTML)
	return err
}
