// +build !nocache

package templatemap

import (
	"errors"
	"html/template"
)

// TemplateMap a map of templates stored
type TemplateMap map[string]*template.Template

// Create creates a map of templates
func Create() TemplateMap {
	return make(TemplateMap)
}

// SetTemplate set a template in the map
func (m TemplateMap) SetTemplate(id string, f func() (*template.Template, error)) error {
	// Greedy evaluation
	tmpl, err := f()
	if err != nil {
		return err
	}
	m[id] = tmpl

	return nil
}

// GetTemplate get a template from the map
func (m TemplateMap) GetTemplate(id string) (*template.Template, error) {
	if m[id] == nil {
		return nil, errors.New("template " + id + " does not exist")
	}
	return m[id], nil
}
