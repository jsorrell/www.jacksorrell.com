package data

import (
	"net/http"
	"path/filepath"
)

// WebDir fulfills http.FileSystem, with assets taken from Assets
type WebDir struct {
	base string
}

// WebPublic is the WebDir containing all public assets.
var WebPublic = &WebDir{"public/"}

// Favicons is the WebDir containing favicons.
var Favicons = &WebDir{"public/fav/"}

// Templates is the WebDir containing all the app's templates.
var Templates = &WebDir{"templates/"}

// Open returns the file descriptor of the file with the name given.
func (d *WebDir) Open(name string) (http.File, error) {
	return Assets.Open(filepath.Join(d.base, name))
}
