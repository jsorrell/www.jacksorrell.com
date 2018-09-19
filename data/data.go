package data

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var WebPublic = &WebDir{"public/"}
var Favicons = &WebDir{"public/fav/"}
var Templates = &WebDir{"templates/"}

type WebDir struct {
	base string
}

func (d *WebDir) Open(name string) (http.File, error) {
	return Assets.Open(filepath.Join(d.base, name))
}

func ReadFileToString(templateFile io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(templateFile)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
