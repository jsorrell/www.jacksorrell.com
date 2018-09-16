// +build nocache

package templatemap

import (
	"errors"
	"html/template"
)

// TemplateMap a map of templates stored
type TemplateMap map[string]func() (*template.Template, error)

// Create creates a map of templates
func Create() TemplateMap {
	return make(TemplateMap)
}

// SetTemplate set a template in the map
func (m TemplateMap) SetTemplate(id string, f func() (*template.Template, error)) error {
	m[id] = f
	return nil
}

// GetTemplate get a template from the map
func (m TemplateMap) GetTemplate(id string) (*template.Template, error) {
	if m[id] == nil {
		return nil, errors.New("template " + id + " does not exist")
	}

	// No cache, so must generate
	tmpl, err := m[id]()
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
