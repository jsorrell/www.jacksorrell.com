// +build !dev

package templates

import (
	"bytes"
	"html/template"

	"github.com/jsorrell/www.jacksorrell.com/log"
	myio "github.com/jsorrell/www.jacksorrell.com/utils/io"
	"github.com/jsorrell/www.jacksorrell.com/web/pages"
)

type _staticTemplateInfo struct {
	tmpl         *staticTemplate
	templateName string
	argCreator   func() interface{}
}

/* Template Group */

type templateGroup struct {
	t       *template.Template
	statics []_staticTemplateInfo
}

func newTemplateGroup(funcs template.FuncMap) *templateGroup {
	return &templateGroup{template.New("").Funcs(funcs), []_staticTemplateInfo{}}
}

func (g *TemplateGroup) getTemplate() *template.Template {
	return g.t
}

func (g *TemplateGroup) init() {
	for _, static := range g.statics {
		buf := bytes.NewBuffer(make([]byte, 0, 5000)) // This stores data, so don't reuse this buf
		if err := g.t.ExecuteTemplate(buf, static.templateName, createArg(static.argCreator)); err != nil {
			log.Fatal(err)
		}
		static.tmpl.compiledTemplate = buf.Bytes()
		static.tmpl.etag = pages.GenEtag(static.tmpl.compiledTemplate)
	}
	g.statics = []_staticTemplateInfo{} // Don't need this anymore so might as well garbage collect
}

/* Static Template */

type staticTemplate struct {
	compiledTemplate []byte
	etag             string
}

func (g *TemplateGroup) newStaticTemplate(templateName, fileName string, createArgs func() interface{}) *staticTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	tmpl := &staticTemplate{[]byte{}, ""}
	g.statics = append(g.statics, _staticTemplateInfo{tmpl, templateName, createArgs})
	return tmpl
}

func (tmpl *StaticTemplate) getReader() (myio.ReadSeekerCloser, string) {
	return myio.RSCloseWrapper{ReadSeeker: bytes.NewReader(tmpl.compiledTemplate)}, tmpl.etag
}

/* Dynamic Template */

type dynamicTemplate struct{}

func (g *TemplateGroup) newDynamicTemplate(templateName, fileName string) *dynamicTemplate {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
	return &dynamicTemplate{}
}

/* Included Template */

func (g *TemplateGroup) newIncludedTemplate(templateName, fileName string) {
	if err := addTemplate(&g.t, templateName, fileName); err != nil {
		log.Fatal(err)
	}
}
